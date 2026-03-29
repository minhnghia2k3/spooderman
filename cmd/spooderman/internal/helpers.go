package internal

import (
	"os"
	"path/filepath"

	"github.com/minhnghia2k3/spooderman/pkg"
	"github.com/minhnghia2k3/spooderman/pkg/config"
	"github.com/minhnghia2k3/spooderman/pkg/logger"
)

// GetSpoodermanHome returns the home directory (e.g. /home/usr/.spooderman)
func GetSpoodermanHome() string {
	if home := os.Getenv(config.EnvHome); home != "" {
		return home
	}

	home, _ := os.UserHomeDir()
	return filepath.Join(home, pkg.DefaultSpoodermanHome)
}

// GetConfigPath from spooderman home (e.g. /home/usr/.spooderman/config.json)
func GetConfigPath() string {
	if configPath := os.Getenv(config.EnvConfig); configPath != "" {
		return configPath
	}

	return filepath.Join(GetSpoodermanHome(), "config.json")
}

func LoadConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig(GetConfigPath())
	if err != nil {
		return nil, err
	}

	logger.SetLevelFromString("DEBUG")
	return cfg, err
}
