package core

import "sort"

func DeduplicateNetworks(networks []Network) []Network {

	best := make(map[string]Network)

	for _, n := range networks {

		var key string

		switch n.Type {

		case TypeWiFi:

			if n.SSID == "" {
				continue
			}

			band := FreqToBand(n.Frequency)

			key = "wifi|" + n.SSID + "|" + band

		case TypeEthernet:

			key = "eth|" + n.Interface

		case TypeVPN:

			key = "vpn|" + n.UUID

		case TypeBluetooth:

			key = "bt|" + n.Interface

		default:
			continue
		}

		existing, ok := best[key]

		if !ok ||
			n.Connected ||
			(!existing.Connected && n.Signal > existing.Signal) {

			best[key] = n
		}
	}

	result := make([]Network, 0, len(best))

	for _, n := range best {
		result = append(result, n)
	}

	return result
}

func SortNetworks(networks []Network) {

	sort.Slice(networks, func(i, j int) bool {

		a := networks[i]
		b := networks[j]

		if a.Connected != b.Connected {
			return a.Connected
		}

		if a.Type != b.Type {

			if a.Type == TypeEthernet {
				return true
			}

			if b.Type == TypeEthernet {
				return false
			}
		}

		return a.Signal > b.Signal
	})
}
