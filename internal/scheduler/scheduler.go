package scheduler

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/MohsenParandvar/reployer/internal/errs"
)

type Scheduler struct {
	interval time.Duration
}

func New(interval time.Duration) *Scheduler {
	return &Scheduler{
		interval: interval,
	}
}

func (s *Scheduler) Run(ctx context.Context, job func(context.Context) error, logger *slog.Logger) error {
	ticker := time.NewTicker(s.interval)

	// first run (immediate)
	if err := job(ctx); err != nil {
		if errors.Is(err, errs.ErrDockerDeamon) {
			return err
		}
	}

	for {
		select {
		case <-ticker.C:
			if err := job(ctx); err != nil {
				if errors.Is(err, errs.ErrDockerDeamon) {
					return err
				}
				logger.Warn(err.Error())
				continue
			}
		case <-ctx.Done():
			ticker.Stop()
			return nil
		}
	}
}
