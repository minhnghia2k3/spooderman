package config

import (
	"os"
	"path/filepath"

	"github.com/minhnghia2k3/spooderman/pkg"
)

// DefaultConfig returns default config for Spooderman
func DefaultConfig() *Config {
	// Get default config
	// Priority: $SPOODERMAN_HOME > ~/.spooderman
	var homePath string
	if path := os.Getenv(EnvHome); path != "" {
		homePath = path
	} else {
		userDir, _ := os.UserHomeDir()
		homePath = filepath.Join(userDir, pkg.DefaultSpoodermanHome)
	}

	workspace := filepath.Join(homePath, pkg.WorkspaceName)

	return &Config{
		Version: CurrentVersion,
		Agents: AgentsConfig{
			Defaults: AgentDefaults{
				Workspace: workspace,
			},
		},
	}
}
