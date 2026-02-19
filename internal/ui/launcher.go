package ui

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/aayushkdev/nmsurf/internal/config"
)

var cfg = config.Load()

func runLauncher(prompt string, input []string, password bool) (string, error) {

	var cmd *exec.Cmd
	var args []string

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

	case "fuzzel":

		args = []string{
			"--dmenu",
			"--prompt", prompt,
		}

		if cfg.Theme != "" {
			args = append(args, "--config", cfg.Theme)
		}

		cmd = exec.Command("fuzzel", args...)

	case "walker":
		fmt.Println("Using walker launcher")

		args = []string{
			"--dmenu",
			"--placeholder", prompt,
		}

		if cfg.Theme != "" {
			args = append(args, "--theme", cfg.Theme)
		}

		cmd = exec.Command("walker", args...)

	default: // wofi

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

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {

		fmt.Printf("Launcher '%s' failed: %v\n", cfg.Launcher, err)

		if stderr.Len() > 0 {
			fmt.Printf("stderr: %s\n", stderr.String())
		}

		return "", err
	}

	return strings.TrimSpace(stdout.String()), nil
}
