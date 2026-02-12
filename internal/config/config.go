package config

type Spec struct {
	File string `yaml:"file"`
}

type Service struct {
	Name     string `yaml:"name"`
	Image    string `yaml:"image"`
	Deployer string `yaml:"deployer"`
	Policy   string `yaml:"update_policy"`
	Spec     Spec   `yaml:"spec"`
}

type Config struct {
	IntervalSeconds int       `yaml:"interval_seconds"`
	Services        []Service `yaml:"services"`
}
