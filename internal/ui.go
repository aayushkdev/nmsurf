package internal

import (
	"bytes"
	"os/exec"
	"strings"
)

func ShowMenu(options []string, prompt string) (string, error) {

	display := make([]string, len(options))

	for i, opt := range options {

		if parts := strings.SplitN(opt, "|", 2); len(parts) == 2 {
			display[i] = parts[1]
		} else {
			display[i] = opt
		}
	}

	cmd := exec.Command("wofi", "--dmenu", "--prompt", prompt)
	cmd.Stdin = strings.NewReader(strings.Join(display, "\n"))

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {

		if exitErr, ok := err.(*exec.ExitError); ok &&
			exitErr.ExitCode() == 10 {
			return "", nil
		}

		return "", err
	}

	selected := strings.TrimSpace(out.String())

	for i, d := range display {
		if strings.TrimSpace(d) == selected {
			return options[i], nil
		}
	}

	return "", nil
}

func PromptPassword(ssid string) (string, error) {

	cmd := exec.Command(
		"wofi",
		"--dmenu",
		"--prompt", "Password for "+ssid,
		"--password",
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
