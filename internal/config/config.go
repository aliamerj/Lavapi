package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BaseURL string `json:"base_url"`
}

func ReadConfig(path string) (*Config, error) {
	var conf Config
	if _, err := os.Stat(path); err != nil {
		return nil, nil
	}

	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config.json: %w", err)
	}
	if err := json.Unmarshal(configFile, &conf); err != nil {
		return nil, fmt.Errorf("Invalid config.json: %w", err)
	}
	return &conf, nil
}
