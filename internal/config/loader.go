package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(configPath string) (*Config, error) {
	var conf Config
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
