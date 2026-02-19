package app

import (
	"strings"
	"sync"
	"time"

	"github.com/aayushkdev/nmsurf/internal/core"
	"github.com/aayushkdev/nmsurf/internal/ui"
)

type Controller struct {

	// Registered providers
	providers []core.Provider

	// Tracks busy operations
	busy map[string]bool

	// Cached network list
	cached []core.Network

	mutex sync.RWMutex
}

func NewController(p []core.Provider) *Controller {

	c := &Controller{
		providers: p,
		busy:      make(map[string]bool),
	}

	// Initial scan
	c.refresh()

	// Background scanner
	go c.scanner()

	return c
}

func (c *Controller) scanner() {

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.refresh()
	}
}

func (c *Controller) refresh() {

	var networks []core.Network

	for _, provider := range c.providers {

		n, err := provider.Scan()
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

		if len(networks) == 0 {

			time.Sleep(200 * time.Millisecond)

			continue
		}

		options := make([]string, len(networks))

		for i := range networks {

			id := networks[i].UniqueID()

			busy := c.busy[id]

			options[i] =
				id + "|" +
					ui.FormatNetwork(networks[i], busy)
		}

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
