package ui

import (
	"fmt"
	"strings"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func FormatNetwork(n core.Network, busy bool) string {

	var parts []string

	parts = append(parts, networkIcon(n))

	if name := n.DisplayName(); name != "" {
		parts = append(parts, name)
	}

	if n.Type == core.TypeWiFi {
		if band := core.FreqToBand(n.Frequency); band != "" {
			parts = append(parts, band)
		}
	}

	if icon := statusIcon(n, busy); icon != "" {
		parts = append(parts, icon)
	}

	return strings.Join(parts, " ")
}

func FormatNetworkMenu(n core.Network) []string {

	if n.Connected {

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

func FormatNetworkDetails(n core.Network) []string {

	sec := n.Security
	if sec == "" {
		sec = "None"
	}

	lines := []string{
		"Type: " + networkTypeName(n.Type),
	}

	if n.SSID != "" {
		lines = append(lines, "SSID: "+n.SSID)
	}

	if n.BSSID != "" {
		lines = append(lines, "BSSID: "+n.BSSID)
	}

	if n.Interface != "" {
		lines = append(lines, "Interface: "+n.Interface)
	}

	if n.UUID != "" {
		lines = append(lines, "UUID: "+n.UUID)
	}

	if sig := formatSignal(n.Signal); sig != "" {
		lines = append(lines, "Signal: "+sig)
	}

	if band := core.FreqToBand(n.Frequency); band != "" {
		lines = append(lines, "Band: "+band)
	}

	lines = append(lines, "Security: "+sec)

	if n.Channel > 0 {
		lines = append(lines, fmt.Sprintf("Channel: %d", n.Channel))
	}

	lines = append(lines, "", "Back")

	return lines
}
