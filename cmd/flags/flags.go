package flags

import "flag"

type CLIFlags struct {
	ConfigPath string
	Daemon     bool
	Update     bool
	Service    string
	Tag        string
}

func ParseFlags() CLIFlags {
	var f CLIFlags
	flag.StringVar(&f.ConfigPath, "config", "reployer.yml", "Configuration YAML file path.")
	flag.BoolVar(&f.Daemon, "daemon", false, "Run in Daemon mode (for automatic image update).")
	flag.BoolVar(&f.Update, "update", false, "Update a Service.")
	flag.StringVar(&f.Service, "service", "", "Service Name in config file.")
	flag.StringVar(&f.Tag, "tag", "", "Change the tag of image.")

	flag.Parse()
	return f
}
