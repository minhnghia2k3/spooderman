package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	env "github.com/caarlos0/env/v11"
	"github.com/minhnghia2k3/spooderman/pkg"
	"github.com/minhnghia2k3/spooderman/pkg/logger"
)

var CurrentVersion = 1

type Config struct {
	Version int          `json:"version"`
	Agents  AgentsConfig `json:"agents"`
}

func LoadConfig(path string) (*Config, error) {
	logger.Debugf("loading config from %s\n", path)

	// TODO: update resolver

	// 1. read the config path
	// if not found -> load default config
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.WarnF("config file not found, trying to use default config", map[string]any{"path": path})
			return DefaultConfig(), nil
		}
		logger.Errorf("failed to read config file: %v", err)
		return nil, err
	}

	// 2. Parse & detect config version from config data
	// Detect malware data?
	var versionInfo struct {
		Version int `json:"version"`
	}

	if err := json.Unmarshal(data, &versionInfo); err != nil {
		return nil, fmt.Errorf("failed to detect version info: %w", err)
	}

	if len(data) <= 10 {
		logger.Warnf("content [%s]", string(data))
		return DefaultConfig(), nil
	}

	// 3. Load config from detected version (support migrate version if needed)
	var cfg *Config
	switch versionInfo.Version {
	case 0:
		return nil, nil
	case CurrentVersion:
		cfg, err = loadConfig(data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported config version: %d", versionInfo.Version)
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// ensure workspace default if not set (e.g. ~/.spooderman/workspace)
	if cfg.Agents.Defaults.Workspace == "" {
		homePath, _ := os.UserHomeDir()
		if spoodermanHome := os.Getenv(EnvHome); spoodermanHome != "" {
			homePath = spoodermanHome
		} else if homePath == "" {
			homePath = filepath.Join(homePath, pkg.DefaultSpoodermanHome)
		}

		cfg.Agents.Defaults.Workspace = filepath.Join(homePath, pkg.WorkspaceName)
	}
	return cfg, nil
}

func (c *Config) WorkspacePath() string {
	return expandHome(c.Agents.Defaults.Workspace)
}

type AgentsConfig struct {
	Defaults AgentDefaults `json:"defaults"`
	List     []AgentConfig `json:"list,omitempty"`
}

type AgentDefaults struct {
	Workspace string `json:"workspace" env:"SPOODERMAN_DEFAULTS_WORKSPACE"`
	ModelName string `json:"model_name" env:"SPOODERMAN_DEFAULTS_MODEL_NAME"`
}

func (d *AgentDefaults) GetModelName() string {
	return d.ModelName
}

type AgentConfig struct {
	ID string `json:"id"`
}

// expandHome replaces '~' to user dir path
// e.g ~/.spooderman/workspace -> /home/usr/.spooderman/workspace
func expandHome(path string) string {
	if path == "" {
		return path
	}

	if path[0] == '~' {
		home, _ := os.UserHomeDir()
		if len(path) > 1 && path[1] == '/' {
			path = home + path[1:]
		}
		return home
	}

	return path
}
