package internal

import (
	"strings"
	"sync"
	"time"
)

type Controller struct {
	busy   map[string]bool
	cached []Network
	mutex  sync.RWMutex
}

func NewController() *Controller {

	c := &Controller{
		busy: make(map[string]bool),
	}

	networks, err := ScanNetworks()
	if err == nil {
		networks = DeduplicateNetworks(networks)
		SortNetworks(networks)
		c.cached = networks
	}

	go c.scanner()

	return c
}

func (c *Controller) scanner() {

	for {

		networks, err := ScanNetworks()
		if err == nil {

			networks = DeduplicateNetworks(networks)
			SortNetworks(networks)

			c.mutex.Lock()
			c.cached = networks
			c.mutex.Unlock()
		}

		time.Sleep(2 * time.Second)
	}
}

func (c *Controller) Run() error {

	for {

		c.mutex.RLock()

		networks := make([]Network, len(c.cached))
		copy(networks, c.cached)
		c.mutex.RUnlock()

		if len(networks) == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		options := make([]string, len(networks))

		for i := range networks {
			busy := c.busy[networks[i].BSSID]
			options[i] = networks[i].BSSID + "|" + FormatNetwork(networks[i], busy)
		}

		choice, err := ShowMenu(options, "Networks")
		if err != nil {
			return err
		}

		if choice == "" {
			if len(c.busy) > 0 {
				continue
			}
			return nil
		}

		parts := strings.SplitN(choice, "|", 2)
		if len(parts) != 2 {
			continue
		}

		bssid := parts[0]

		for i := range networks {
			if networks[i].BSSID == bssid {
				c.networkMenu(&networks[i])
				break
			}
		}
	}
}

func (c *Controller) networkMenu(n *Network) {
	for {
		choice, err := ShowMenu(
			FormatNetworkMenu(*n),
			n.SSID,
		)
		if err != nil {
			return
		}

		if choice == "" {
			continue
		}

		switch choice {

		case "Connect":
			c.busy[n.BSSID] = true

			go func(net *Network) {
				defer delete(c.busy, net.BSSID)

				if net.Secured && !net.Saved {
					pass, _ := PromptPassword(net.SSID)

					if pass != "" {
						Connect(net.BSSID, pass)
					}
				} else {
					Connect(net.BSSID, "")
				}
			}(n)

			return

		case "Disconnect":
			c.busy[n.BSSID] = true

			go func(net *Network) {
				defer delete(c.busy, net.BSSID)
				Disconnect(net.SSID)
			}(n)

			return

		case "Forget":
			c.busy[n.BSSID] = true

			go func(net *Network) {
				defer delete(c.busy, net.BSSID)
				Forget(net.SSID)
			}(n)

			return

		case "Details":
			_, err := ShowMenu(
				FormatNetworkDetails(*n),
				"Details",
			)
			if err != nil {
				return
			}

		case "Back":
			return
		}
	}
}
