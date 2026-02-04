package flags

import "flag"

type CLIFlags struct {
	ConfigPath string
	Daemon     bool
}

func ParseFlags() CLIFlags {
	var f CLIFlags
	flag.StringVar(&f.ConfigPath, "config", "reployer.yml", "Configuration YAML file path.")
	flag.BoolVar(&f.Daemon, "daemon", false, "Run in Daemon mode (for automatic image update)")

	flag.Parse()
	return f
}
