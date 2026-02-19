package internal

import (
	"os/exec"
	"strconv"
	"strings"
)

func ScanNetworks() ([]Network, error) {

	savedMap, err := getSavedConnections()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-e", "no",
		"-f",
		"IN-USE,SSID,SIGNAL,SECURITY,FREQ,CHAN,DEVICE,ACTIVE,BSSID",
		"device",
		"wifi",
		"list",
		"--rescan", "no",
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

		secured := security != ""

		networks = append(networks, Network{
			InUse:    inUse,
			SSID:     ssid,
			BSSID:    bssid,
			Signal:   signal,
			Security: security,
			Freq:     freq,
			Channel:  channel,
			Saved:    savedMap[ssid],
			Secured:  secured,
		})
	}

	return networks, nil
}

func getSavedConnections() (map[string]bool, error) {

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-f",
		"NAME",
		"connection",
		"show",
	)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	saved := make(map[string]bool)

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		if line != "" {
			saved[line] = true
		}
	}

	return saved, nil
}

func Connect(bssid string, password string) error {

	if password == "" {
		return exec.Command(
			"nmcli",
			"device",
			"wifi",
			"connect",
			bssid,
		).Run()
	}

	return exec.Command(
		"nmcli",
		"device",
		"wifi",
		"connect",
		bssid,
		"password",
		password,
	).Run()
}

func Disconnect(ssid string) error {

	return exec.Command(
		"nmcli",
		"connection",
		"down",
		ssid,
	).Run()
}

func Forget(ssid string) error {

	return exec.Command(
		"nmcli",
		"connection",
		"delete",
		ssid,
	).Run()
}
