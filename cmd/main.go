package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/MohsenParandvar/reployer/cmd/flags"
	"github.com/MohsenParandvar/reployer/internal/config"
	"github.com/MohsenParandvar/reployer/internal/engine"
	"github.com/MohsenParandvar/reployer/internal/scheduler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := slog.Default()

	cliFlags := flags.ParseFlags()
	configs, err := config.Load(cliFlags.ConfigPath)
	if err != nil {
		logger.Error("Can'not load config", "error", err)
		os.Exit(1)
	}

	eng := engine.New(configs, logger)

	if cliFlags.Daemon {
		if cliFlags.Update {
			logger.Error("Dont use the -daemon and -update in the same time")
			os.Exit(1)
		}

		sch := scheduler.New(time.Duration(configs.IntervalSeconds) * time.Second)

		logger.Info("Daemon Mode Started...")
		err = sch.Run(ctx, func(ctx context.Context) error {
			return eng.Check(ctx)
		}, logger)

		if err != nil {
			logger.Error("Daemon engine returns error", "error", err)
		}
	}

	if cliFlags.Update {
		if cliFlags.Service == "" {
			logger.Info("Set the -service for update")
			os.Exit(1)
		}

		eng.ManualDeploy(ctx, cliFlags.Service, cliFlags.Tag)
	}

}
