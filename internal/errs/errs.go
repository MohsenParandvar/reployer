package errs

import "errors"

var (
	ErrDockerDeamon       = errors.New("docker daemon is not running/installed or image was not found")
	ErrEmptyYamlFile      = errors.New("Yaml file is empty")
	ErrServiceNotFound    = errors.New("Service not found")
	ErrServiceMappingNode = errors.New("Service is not a mapping node")
)
