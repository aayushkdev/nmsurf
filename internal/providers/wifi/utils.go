package wifi

import (
	"os/exec"
	"strings"
)

func getSavedSSIDs() (map[string]bool, error) {

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

	for _, ssid := range lines {

		if ssid == "" {
			continue
		}

		saved[ssid] = true
	}

	return saved, nil
}
