package errs

import "errors"

var (
	ErrDockerDeamon = errors.New("docker daemon is not running/installed")
)
