package docker

import (
	"context"
	"errors"
	"log/slog"
	"os/exec"

	"github.com/MohsenParandvar/reployer/internal/config"
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

func DeployComposeService(ctx context.Context, service config.Service, logger *slog.Logger) error {
	composeServices, err := GetComposeServices(service.Spec.File)

	if err != nil {
		return err
	}

	if csName, csExists := composeServices[service.Name]; csExists {
		digestMatch, err := CompareDigest(ctx, csName)
		if err != nil {
			return err
		}

		if !digestMatch {
			logger.Info("new image found for", "service", service.Name, "update_policy", service.Policy)

			if service.Policy == "update" {
				logger.Info("Start Deploying", "service", service.Name)

				if err := PullComposeImage(service.Spec.File, service.Name); err != nil {
					return err
				}

				logger.Info("Image pulled from remote registry", "image", csName, "service", service.Name)
				logger.Info("Restarting container", "service", service.Name)

				if err := RestartContainer(service.Spec.File, service.Name); err != nil {
					return errors.New("container restarting failed")
				}

				logger.Info("Container restarted", "service", service.Name)
			}
		}
	}

	return nil
}
