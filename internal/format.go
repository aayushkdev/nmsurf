package internal

import (
	"fmt"
)

func FreqToBand(freq int) string {

	ghz := float64(freq) / 1000.0

	return fmt.Sprintf("%.1f GHz", ghz)
}

func SignalIcon(signal int, saved bool, secured bool) string {

	locked := secured && !saved

	switch {
	case signal >= 80:
		if locked {
			return "󰤪 "
		}
		return "󰤨 "

	case signal >= 60:
		if locked {
			return "󰤧 "
		}
		return "󰤥 "

	case signal >= 40:
		if locked {
			return "󰤤 "
		}
		return "󰤢 "

	case signal >= 20:
		if locked {
			return "󰤡 "
		}
		return "󰤟 "

	default:
		return "󰤫 "
	}
}

func FormatNetwork(n Network) string {

	check := "  "
	if n.InUse {
		check = "󰄬 "
	}

	return fmt.Sprintf(
		"%s  %s %s  %s",
		SignalIcon(n.Signal, n.Saved, n.Secured),
		n.SSID,
		FreqToBand(n.Freq),
		check,
	)
}

func FormatNetworkMenu(n Network) []string {

	if n.InUse {
		return []string{
			"Disconnect",
			"Forget",
			"Details",
			"Back",
		}
	}

	return []string{
		"Connect",
		"Forget",
		"Details",
		"Back",
	}
}

func FormatNetworkDetails(n Network) []string {

	sec := n.Security
	if sec == "" {
		sec = "None"
	}

	return []string{
		"SSID: " + n.SSID,
		"BSSID: " + n.BSSID,
		fmt.Sprintf("Signal: %d%%", n.Signal),
		"Band: " + FreqToBand(n.Freq),
		"Security: " + sec,
		fmt.Sprintf("Channel: %d", n.Channel),
		"",
		"Back",
	}
}
