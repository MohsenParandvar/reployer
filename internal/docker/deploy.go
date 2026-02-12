package docker

import (
	"os/exec"
)

func ComposeDeploy(composeFile string, serviceName string) error {
	command := exec.Command("docker-compose", "-f", composeFile, "up", "-d", serviceName)

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}
