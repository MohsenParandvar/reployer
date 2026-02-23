package errs

import "errors"

var (
	ErrDockerDeamon            = errors.New("docker daemon is not running/installed or image was not found")
	ErrEmptyYamlFile           = errors.New("Yaml file is empty")
	ErrKeyNotFound             = errors.New("Key not found")
	ErrServiceNotFound         = errors.New("Service not found")
	ErrNotMappingNode          = errors.New("Can't find a mapping node")
	ErrInvalidImageNode        = errors.New("Invalid image node")
	ErrDigestImageNotSupported = errors.New("Image digest not supported")
	ErrRollBackFailed          = errors.New("Rollback failed")
)
