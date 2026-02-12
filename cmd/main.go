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
		sch := scheduler.New(time.Duration(configs.IntervalSeconds) * time.Second)

		logger.Info("Daemon Starting...")
		err = sch.Run(ctx, func(ctx context.Context) error {
			return eng.Check(ctx)
		})

		if err != nil {
			logger.Error("", "error", err)
		}
	}

}
