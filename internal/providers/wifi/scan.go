package wifi

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func (p *Provider) Scan(hard bool) ([]core.Network, error) {

	saved, err := getSavedSSIDs()
	if err != nil {
		return nil, err
	}

	rescanArg := "no"
	if hard {
		rescanArg = "yes"
	}

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-e", "no",
		"-f",
		"IN-USE,SSID,SIGNAL,SECURITY,FREQ,CHAN,DEVICE,BSSID",
		"device",
		"wifi",
		"list",
		"--rescan", rescanArg,
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	networks := make([]core.Network, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")

		if len(fields) < 8 {
			continue
		}

		inUse := fields[0] == "*"

		ssid := fields[1]

		signal, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		security := fields[3]

		freqStr := strings.TrimSuffix(fields[4], " MHz")
		freq, err := strconv.Atoi(freqStr)
		if err != nil {
			continue
		}

		channel, err := strconv.Atoi(fields[5])
		if err != nil {
			continue
		}

		device := fields[6]
		bssid := fields[7]

		networks = append(networks, core.Network{
			Type: core.TypeWiFi,

			SSID:      ssid,
			BSSID:     bssid,
			Interface: device,

			Signal: signal,

			Security: security,
			Secured:  security != "",
			Saved:    saved[ssid],

			Frequency: freq,
			Channel:   channel,

			Connected: inUse,
		})
	}
	return networks, nil
}
