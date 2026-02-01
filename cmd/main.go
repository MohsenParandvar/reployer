package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MohsenParandvar/reployer/cmd/flags"
	"github.com/MohsenParandvar/reployer/internal/config"
	"github.com/MohsenParandvar/reployer/internal/engine"
	"github.com/MohsenParandvar/reployer/internal/scheduler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cliFlags := flags.ParseFlags()
	configs, err := config.Load(cliFlags.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	eng := engine.New(configs)

	sch := scheduler.New(time.Duration(configs.IntervalSeconds) * time.Second)

	err = sch.Run(ctx, func(ctx context.Context) error {
		return eng.Check(ctx)
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(configs.IntervalSeconds)

}
