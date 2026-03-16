package ui

import (
	"fmt"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func signalIcon(signal int, saved bool, secured bool) string {

	locked := secured && !saved

	switch {

	case signal >= 80:
		if locked {
			return "󰤪  "
		}
		return "󰤨  "

	case signal >= 60:
		if locked {
			return "󰤧  "
		}
		return "󰤥  "

	case signal >= 40:
		if locked {
			return "󰤤  "
		}
		return "󰤢  "

	case signal >= 20:
		if locked {
			return "󰤡  "
		}
		return "󰤟  "
	}

	return "󰤫  "
}

func networkIcon(n core.Network) string {

	switch n.Type {

	case core.TypeWiFi:

		return signalIcon(
			n.Signal,
			n.Saved,
			n.Secured,
		)

	case core.TypeEthernet:

		if n.Connected {
			return "󰈀  "
		}

		return "󰈂  "

	case core.TypeVPN:

		if n.Connected {
			return "󰦝  "
		}

		return "󰦞  "

	case core.TypeBluetooth:

		if n.Connected {
			return "󰂯  "
		}

		return "󰂲  "
	}

	return "󰈂  "
}

func statusIcon(n core.Network) string {

	if n.Connected {
		return " 󰄬"
	}

	return ""
}

func networkTypeName(t core.NetworkType) string {

	switch t {

	case core.TypeWiFi:
		return "WiFi"

	case core.TypeEthernet:
		return "Ethernet"

	case core.TypeVPN:
		return "VPN"

	case core.TypeBluetooth:
		return "Bluetooth"
	}

	return "Unknown"
}

func formatSignal(signal int) string {

	if signal <= 0 {
		return ""
	}

	return fmt.Sprintf("%d%%", signal)
}
