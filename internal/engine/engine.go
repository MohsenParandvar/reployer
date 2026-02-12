package engine

import (
	"context"
	"errors"
	"log/slog"

	"github.com/MohsenParandvar/reployer/internal/config"
	"github.com/MohsenParandvar/reployer/internal/docker"
)

type Engine struct {
	cfg   *config.Config
	state map[string]string
	log   *slog.Logger
}

func New(configs *config.Config, logger *slog.Logger) *Engine {
	return &Engine{
		cfg:   configs,
		state: make(map[string]string),
		log:   logger,
	}
}

func (e *Engine) Check(ctx context.Context) error {
	for _, service := range e.cfg.Services {
		switch service.Deployer {
		case "compose":
			composeServices, err := docker.GetComposeServices(service.Spec.File)

			if err != nil {
				return err
			}

			if csName, csExists := composeServices[service.Name]; csExists {
				digestMatch, err := docker.CompareDigest(ctx, csName)
				if err != nil {
					return err
				}

				if !digestMatch {
					e.log.Info("new image found for", "service", service.Name, "update_policy", service.Policy)

					if service.Policy == "update" {
						e.log.Info("Start Deploying", "service", service.Name)

						if err := docker.PullComposeImage(service.Spec.File, service.Name); err != nil {
							return errors.New("can not pull docker image")
						}

						e.log.Info("Image pulled from remote registry", "image", csName, "service", service.Name)
						e.log.Info("Restarting container", "service", service.Name)

						if err := docker.RestartContainer(service.Spec.File, service.Name); err != nil {
							return errors.New("container restarting failed")
						}

						e.log.Info("Container restarted", "service", service.Name)
					}
				}
			}
		}
	}

	return nil
}
