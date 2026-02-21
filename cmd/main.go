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
		logger.Error(err.Error())
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
			logger.Error(err.Error())
		}
	}

	if cliFlags.Update {
		if cliFlags.Service == "" {
			logger.Info("Set the -service for update")
			os.Exit(1)
		}

		if err := eng.ManualDeploy(ctx, cliFlags.Service, cliFlags.Tag); err != nil {
			logger.Error(err.Error())
		}
	}

	if !cliFlags.Daemon && !cliFlags.Update {
		logger.Info("Use -daemon for starting daemon mode and -update for update container mannually")
	}
}
