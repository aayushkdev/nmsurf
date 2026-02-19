package ui

import (
	"bytes"
	"os/exec"
	"strings"
)

const (
	RescanID = "__rescan__"
)

func ShowMenu(options []string, prompt string) (string, error) {

	display := make([]string, len(options))

	for i, opt := range options {

		parts := strings.SplitN(opt, "|", 2)

		if len(parts) == 2 {
			display[i] = parts[1]
		} else {
			display[i] = opt
		}
	}

	cmd := exec.Command(
		"wofi",
		"--dmenu",
		"--prompt", prompt,
	)

	cmd.Stdin = strings.NewReader(
		strings.Join(display, "\n"),
	)

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {

		if exitErr, ok := err.(*exec.ExitError); ok &&
			exitErr.ExitCode() == 10 {

			return "", nil
		}

		return "", err
	}

	selected := strings.TrimSpace(out.String())

	for i, d := range display {

		if d == selected {

			return options[i], nil
		}
	}

	return "", nil
}

func PromptPassword(name string) (string, error) {

	cmd := exec.Command(
		"wofi",
		"--dmenu",
		"--prompt", "Password for "+name,
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
