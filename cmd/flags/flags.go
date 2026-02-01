package flags

import "flag"

type CLIFlags struct {
	ConfigPath string
}

func ParseFlags() CLIFlags {
	var f CLIFlags
	flag.StringVar(&f.ConfigPath, "config", "reployer.yml", "Configuration YAML file path.")

	flag.Parse()
	return f
}
