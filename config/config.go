package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
