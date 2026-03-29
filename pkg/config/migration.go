package config

import "encoding/json"

func loadConfig(data []byte) (*Config, error) {
	cfg := DefaultConfig()

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
