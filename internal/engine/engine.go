package engine

import (
	"context"
	"fmt"

	"github.com/MohsenParandvar/reployer/internal/config"
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
	fmt.Println("hello")
	return nil
}
