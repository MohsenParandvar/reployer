package config

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func (conf *Config) LoadConfig(configPath string) *Config {
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
