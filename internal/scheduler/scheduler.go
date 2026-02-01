package scheduler

import (
	"context"
	"time"
)

type Scheduler struct {
	interval time.Duration
}

func New(interval time.Duration) *Scheduler {
	return &Scheduler{
		interval: interval,
	}
}

func (s *Scheduler) Run(ctx context.Context, job func(context.Context) error) error {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	// first run (immediate)
	if err := job(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			if err := job(ctx); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}
