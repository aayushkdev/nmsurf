package internal

import (
	"os/exec"
	"strconv"
	"strings"
)

func ScanNetworks() ([]Network, error) {

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-e", "no",
		"-f",
		"IN-USE,SSID,SIGNAL,SECURITY,FREQ,CHAN,DEVICE,ACTIVE,BSSID",
		"device",
		"wifi",
		"list",
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	var networks []Network

	for _, line := range lines {

		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")

		if len(fields) < 14 {
			continue
		}

		inUse := fields[0] == "*"
		ssid := fields[1]

		signal, _ := strconv.Atoi(fields[2])

		security := fields[3]

		freqStr := strings.TrimSuffix(fields[4], " MHz")
		freq, _ := strconv.Atoi(freqStr)

		channel, _ := strconv.Atoi(fields[5])

		bssid := strings.Join(fields[len(fields)-6:], ":")

		networks = append(networks, Network{
			InUse:    inUse,
			SSID:     ssid,
			BSSID:    bssid,
			Signal:   signal,
			Security: security,
			Freq:     freq,
			Channel:  channel,
		})
	}

	return networks, nil
}

func Connect(bssid string) error {

	cmd := exec.Command(
		"nmcli",
		"device",
		"wifi",
		"connect",
		bssid,
	)

	return cmd.Run()
}

func Disconnect(ssid string) error {

	cmd := exec.Command(
		"nmcli",
		"connection",
		"down",
		ssid,
	)

	return cmd.Run()
}

func Forget(ssid string) error {

	cmd := exec.Command(
		"nmcli",
		"connection",
		"delete",
		ssid,
	)

	return cmd.Run()
}
