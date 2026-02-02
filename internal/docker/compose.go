package docker

import (
	"os"

	"github.com/goccy/go-yaml"
)

type ComposeService struct {
	Image string `yaml:"image"`
}

type ComposeFile struct {
	Services map[string]ComposeService `yaml:"services"`
}

func GetComposeServices(composeFilePath string) (map[string]string, error) {
	file, err := os.ReadFile(composeFilePath)

	if err != nil {
		return nil, err
	}

	var compose ComposeFile
	if err := yaml.Unmarshal(file, &compose); err != nil {
		return nil, err
	}

	images := make(map[string]string)
	for service, config := range compose.Services {
		images[service] = config.Image
	}

	return images, nil
}
