package internal

import "sort"

type Network struct {
	InUse      bool
	SSID       string
	Signal     int
	Security   string
	Freq       int
	BSSID      string
	Channel    int
	Saved      bool
	Secured    bool
	Connecting bool
}

func DeduplicateNetworks(networks []Network) []Network {

	best := make(map[string]Network)

	for _, n := range networks {

		if n.SSID == "" {
			continue
		}

		key := n.SSID + "|" + FreqToBand(n.Freq)

		existing, ok := best[key]

		if !ok || n.Signal > existing.Signal || n.InUse {
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

		if networks[i].InUse != networks[j].InUse {
			return networks[i].InUse
		}

		return networks[i].Signal > networks[j].Signal
	})
}
