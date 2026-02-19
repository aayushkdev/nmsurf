package internal

import "strings"

type Controller struct {
	busy map[string]bool
}

func NewController() *Controller {
	return &Controller{
		busy: make(map[string]bool),
	}
}

func (c *Controller) Run() error {
	for {
		networks, err := ScanNetworks()
		if err != nil {
			return err
		}

		networks = DeduplicateNetworks(networks)
		SortNetworks(networks)

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

		var selected *Network

		for i := range networks {
			if networks[i].BSSID == bssid {
				selected = &networks[i]
				break
			}
		}

		if selected != nil {
			c.networkMenu(selected)
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
