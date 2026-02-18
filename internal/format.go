package internal

import (
	"fmt"
)

func FreqToBand(freq int) string {

	ghz := float64(freq) / 1000.0

	return fmt.Sprintf("%.1f GHz", ghz)
}

func SignalIcon(signal int, saved bool, connected bool) string {

	if connected && saved {
		switch {
		case signal >= 80:
			return "󰤨 "
		case signal >= 60:
			return "󰤥 "
		case signal >= 40:
			return "󰤢 "
		case signal >= 20:
			return "󰤟 "
		default:
			return "󰤫 "
		}
	}

	switch {
	case signal >= 80:
		return "󰤪 "
	case signal >= 60:
		return "󰤧 "
	case signal >= 40:
		return "󰤤 "
	case signal >= 20:
		return "󰤡 "
	default:
		return "󰤫 "
	}
}

func FormatNetwork(n Network) string {

	check := "  "
	if n.InUse {
		check = "󰄬 "
	}

	lock := ""
	if n.Security != "" {
		lock = "󰌾"
	}

	return fmt.Sprintf(
		"%s%s %s %s  %s",
		check,
		SignalIcon(n.Signal, n.Saved, n.InUse),
		n.SSID,
		FreqToBand(n.Freq),
		lock,
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
