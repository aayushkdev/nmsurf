package ui

import (
	"strings"
)

const (
	RescanID     = "__rescan__"
	WifiToggleID = "__wifi_toggle__"
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

	selected, err := runLauncher(prompt, display, false)
	if err != nil {
		return "", err
	}

	for i, d := range display {
		if d == selected {
			return options[i], nil
		}
	}

	return "", nil
}

func PromptPassword(name string) (string, error) {

	return runLauncher(
		"Password for "+name,
		nil,
		true,
	)
}
