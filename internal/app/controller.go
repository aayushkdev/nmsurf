package app

import (
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/aayushkdev/nmsurf/internal/core"
	"github.com/aayushkdev/nmsurf/internal/ui"
)

type Controller struct {
	providers  []core.Provider
	busy       map[string]bool
	cached     []core.Network
	mutex      sync.RWMutex
	rescanning bool
	wifiOn     bool
}

func NewController(p []core.Provider) *Controller {

	c := &Controller{
		providers: p,
		busy:      make(map[string]bool),
	}

	// determine initial wifi radio state (nmcli returns "enabled"/"disabled")
	if out, err := exec.Command("nmcli", "radio", "wifi").Output(); err == nil {
		c.wifiOn = strings.TrimSpace(string(out)) == "enabled"
	} else {
		c.wifiOn = true
	}

	// Initial scan
	c.refresh(false)

	go c.scanner()

	return c
}

func (c *Controller) scanner() {

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.refresh(false)
	}
}

func (c *Controller) refresh(hard bool) {

	var networks []core.Network

	for _, provider := range c.providers {

		n, err := provider.Scan(hard)
		if err != nil {
			continue
		}

		networks = append(networks, n...)
	}

	networks = core.DeduplicateNetworks(networks)
	core.SortNetworks(networks)

	c.mutex.Lock()
	c.cached = networks
	c.mutex.Unlock()
}

func (c *Controller) getCached() []core.Network {

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	out := make([]core.Network, len(c.cached))

	copy(out, c.cached)

	return out
}

func (c *Controller) Run() error {

	for {

		networks := c.getCached()

		// top fixed: WiFi toggle
		options := make([]string, 0, len(networks)+2)
		options = append(options, ui.WifiToggleID+"|"+ui.FormatWifiToggle(c.wifiOn))

		for i := range networks {

			id := networks[i].UniqueID()
			busy := c.busy[id]

			options = append(options,
				id+"|"+ui.FormatNetwork(networks[i], busy),
			)
		}

		options = append(options, ui.RescanID+"|󰑐  Rescan")
		choice, err := ui.ShowMenu(options, "Networks")

		if err != nil {
			return err
		}

		if choice == "" {
			if len(c.busy) > 0 {
				continue
			}

			return nil
		}

		id := strings.SplitN(choice, "|", 2)[0]
		switch id {

		case ui.WifiToggleID:
			c.toggleWifi()
			continue

		case ui.RescanID:
			c.rescanning = true
			c.refresh(true)
			c.rescanning = false
			continue
		}

		for i := range networks {

			if networks[i].UniqueID() == id {

				c.networkMenu(&networks[i])

				break
			}
		}
	}
}

func (c *Controller) networkMenu(n *core.Network) {

	for {

		choice, err := ui.ShowMenu(
			ui.FormatNetworkMenu(*n),
			n.DisplayName(),
		)

		if err != nil {
			return
		}

		if choice == "" {
			continue
		}

		switch choice {

		case "Connect":

			c.runAsync(n, func(p core.Provider) {

				if n.Secured && !n.Saved {

					pass, _ :=
						ui.PromptPassword(n.DisplayName())

					if pass != "" {
						p.Connect(*n, pass)
					}

				} else {

					p.Connect(*n, "")
				}
			})

			return

		case "Disconnect":

			c.runAsync(n, func(p core.Provider) {
				p.Disconnect(*n)
			})

			return

		case "Forget":

			c.runAsync(n, func(p core.Provider) {
				p.Forget(*n)
			})

			return

		case "Details":

			ui.ShowMenu(
				ui.FormatNetworkDetails(*n),
				"Details",
			)

		case "Back":
			return
		}
	}
}

func (c *Controller) runAsync(
	n *core.Network,
	fn func(core.Provider),
) {

	id := n.UniqueID()

	c.busy[id] = true

	go func() {

		defer delete(c.busy, id)

		for _, p := range c.providers {
			fn(p)
		}
	}()
}

func (c *Controller) toggleWifi() {
	desired := !c.wifiOn

	var cmd *exec.Cmd
	if desired {
		cmd = exec.Command("nmcli", "radio", "wifi", "on")
	} else {
		cmd = exec.Command("nmcli", "radio", "wifi", "off")
	}

	// attempt to toggle, ignore errors but update local state
	_ = cmd.Run()

	c.wifiOn = desired

	// refresh networks after toggling
	c.refresh(true)
}
