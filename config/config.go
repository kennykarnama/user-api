package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RestPort string `envconfig:"REST_PORT" default:"8080"`
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
