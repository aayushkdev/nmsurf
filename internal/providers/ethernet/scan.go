package ethernet

import (
	"os/exec"
	"strings"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func (p *Provider) Scan(hard bool) ([]core.Network, error) {

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-f",
		"DEVICE,TYPE,STATE,CONNECTION",
		"device",
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	var networks []core.Network

	for _, line := range lines {

		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")

		if len(fields) < 4 {
			continue
		}

		device := fields[0]
		devType := fields[1]
		state := fields[2]

		if devType != "ethernet" {
			continue
		}

		networks = append(networks, core.Network{
			Type:      core.TypeEthernet,
			Interface: device,
			Connected: state == "connected",
		})
	}

	return networks, nil
}
