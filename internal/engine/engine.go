package engine

import (
	"context"
	"log/slog"

	"github.com/MohsenParandvar/reployer/internal/config"
	"github.com/MohsenParandvar/reployer/internal/docker"
	"github.com/MohsenParandvar/reployer/internal/errs"
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
			if err := docker.DeployComposeService(ctx, service, e.log); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Engine) ManualDeploy(ctx context.Context, serviceName string, tag string) error {
	for _, service := range e.cfg.Services {
		if service.Name == serviceName {
			switch service.Deployer {
			case "compose":
				if tag != "" {
					e.log.Info("Changing image tag", "service", serviceName, "tag", tag)

					if err := docker.ChangeServiceTag(service.Spec.File, serviceName, tag); err != nil {
						return err
					}

					e.log.Info("Image tag changed successfully", "service", serviceName, "tag", tag)
				}

				if err := docker.DeployComposeService(ctx, service, e.log); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return errs.ErrServiceNotFound
}
