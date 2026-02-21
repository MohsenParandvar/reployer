package docker

import (
	"os"
	"strings"

	"github.com/MohsenParandvar/reployer/internal/errs"
	"gopkg.in/yaml.v3"
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

func ChangeServiceTag(composeFilePath string, serviceName string, tag string) error {
	file, err := os.ReadFile(composeFilePath)
	if err != nil {
		return err
	}

	var node yaml.Node
	if err := yaml.Unmarshal(file, &node); err != nil {
		return err
	}

	if err := SetServiceImage(&node, serviceName, tag); err != nil {
		return err
	}

	exportFile, err := os.Create(composeFilePath)
	if err != nil {
		return err
	}

	encoder := yaml.NewEncoder(exportFile)
	encoder.SetIndent(2)
	if err := encoder.Encode(&node); err != nil {
		encoder.Close()
		return err
	}
	return encoder.Close()
}

func SetServiceImage(node *yaml.Node, serviceName string, tag string) error {
	if node == nil || len(node.Content) == 0 {
		return errs.ErrEmptyYamlFile
	}

	root := node.Content[0]
	services, err := FindMappingChild(root, "services")
	if err != nil {
		return err
	}

	if services.Kind != yaml.MappingNode {
		return errs.ErrServiceMappingNode
	}

	service, err := FindMappingChild(services, serviceName)
	if err != nil {
		return err
	}

	image, err := FindMappingChild(service, "image")
	if err != nil {
		return err
	}

	splittedImage := strings.Split(image.Value, ":")

	image.Kind = yaml.ScalarNode
	image.Tag = "!!str"
	image.Value = splittedImage[0] + ":" + tag

	return nil
}

func FindMappingChild(node *yaml.Node, key string) (*yaml.Node, error) {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil, errs.ErrServiceMappingNode
	}

	for i := 0; i < len(node.Content); i += 2 {
		k := node.Content[i]
		v := node.Content[i+1]
		if k.Value == key {
			return v, nil
		}
	}

	return nil, errs.ErrServiceNotFound
}
