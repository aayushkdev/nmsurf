package internal

import "strings"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
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
			options[i] = networks[i].BSSID + "|" + FormatNetwork(networks[i])
		}

		choice, err := ShowMenu(options, "Networks")
		if err != nil || choice == "" {
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

		if selected == nil {
			continue
		}

		c.networkMenu(selected)
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

			if n.Secured && !n.Saved {

				pass, err := PromptPassword(n.SSID)
				if err != nil || pass == "" {
					return
				}

				Connect(n.BSSID, pass)

			} else {

				Connect(n.BSSID, "")
			}

			return

		case "Disconnect":
			Disconnect(n.SSID)
			return

		case "Forget":
			Forget(n.SSID)
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
