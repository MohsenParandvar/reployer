package main

import (
	"fmt"

	"github.com/MohsenParandvar/reployer/cmd/flags"
	"github.com/MohsenParandvar/reployer/internal/config"
)

func main() {
	flags := flags.ParseFlags()
	var config config.Config
	config.LoadConfig(flags.ConfigPath)

	fmt.Println(config.IntervalSeconds)

}
