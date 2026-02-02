package engine

import (
	"context"
	"fmt"

	"github.com/MohsenParandvar/reployer/internal/config"
	"github.com/MohsenParandvar/reployer/internal/docker"
)

type Engine struct {
	cfg   *config.Config
	state map[string]string
}

func New(configs *config.Config) *Engine {
	return &Engine{
		cfg:   configs,
		state: make(map[string]string),
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
				fmt.Println(csName)
			}
		}
	}

	return nil
}
