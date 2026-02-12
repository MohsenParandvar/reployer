package docker

import (
	"os/exec"
)

func PullComposeImage(composeFile string, serviceName string) error {
	command := exec.Command("docker-compose", "-f", composeFile, "pull", serviceName)

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func RestartContainer(composeFile string, serviceName string) error {
	command := exec.Command("docker-compose", "-f", composeFile, "up", "-d", serviceName)

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}
