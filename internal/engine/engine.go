package engine

import (
	"context"
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
					e.log.Info("new image found for service", "service", csName)
				}
			}
		}
	}

	return nil
}
