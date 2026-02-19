package ui

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/aayushkdev/nmsurf/internal/config"
)

var cfg = config.Load()

func runLauncher(prompt string, input []string, password bool) (string, error) {

	var args []string
	var cmd *exec.Cmd

	switch cfg.Launcher {

	case "rofi":

		args = []string{
			"-dmenu",
			"-p", prompt,
		}

		if password {
			args = append(args, "-password")
		}

		if cfg.Theme != "" {
			args = append(args, "-theme", cfg.Theme)
		}

		cmd = exec.Command("rofi", args...)

	default:

		args = []string{
			"--dmenu",
			"--prompt", prompt,
		}

		if password {
			args = append(args, "--password")
		}

		if cfg.Theme != "" {
			args = append(args, "--style", cfg.Theme)
		}

		cmd = exec.Command("wofi", args...)
	}

	cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
