package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Launcher string `toml:"launcher"`
	Theme    string `toml:"theme"`
}

var Default = Config{
	Launcher: "wofi",
	Theme:    "",
}

func expand(path string) string {

	if path == "" {
		return ""
	}

	if path[0] == '~' {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[1:])
	}

	return path
}

func Load() Config {

	dir, err := os.UserConfigDir()
	if err != nil {
		return Default
	}

	path := filepath.Join(dir, "nmsurf", "config.toml")

	data, err := os.ReadFile(path)
	if err != nil {
		return Default
	}

	cfg := Default

	if err := toml.Unmarshal(data, &cfg); err != nil {
		fmt.Println(err)
		return Default
	}

	cfg.Theme = expand(cfg.Theme)

	if cfg.Launcher == "" {
		cfg.Launcher = "wofi"
	}
	return cfg
}
