package wifi

import (
	"os/exec"
	"strings"
)

func splitNmcliFields(line string) []string {
	fields := make([]string, 0, 8)
	var current strings.Builder
	escaped := false

	for _, r := range line {
		switch {
		case escaped:
			current.WriteRune(r)
			escaped = false

		case r == '\\':
			escaped = true

		case r == ':':
			fields = append(fields, current.String())
			current.Reset()

		default:
			current.WriteRune(r)
		}
	}

	if escaped {
		current.WriteByte('\\')
	}

	fields = append(fields, current.String())

	return fields
}

func getSavedSSIDs() (map[string]bool, error) {

	cmd := exec.Command(
		"nmcli",
		"-t",
		"-e", "yes",
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

		fields := splitNmcliFields(ssid)
		if len(fields) == 0 || fields[0] == "" {
			continue
		}

		saved[fields[0]] = true
	}

	return saved, nil
}
