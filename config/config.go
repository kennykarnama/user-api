package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	RestPort        string `envconfig:"REST_PORT" default:"8080"`
	ServiceName     string `envconfig:"SERVICE_NAME" default:"user_api"`
	AccessTokenCfg  AccessTokenConfig
	RefreshTokenCfg RefreshTokenConfig
}

type AccessTokenConfig struct {
	Secret     string        `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
	Expiration time.Duration `envconfig:"ACCESS_TOKEN_EXPIRATION" default:"15m"`
}

type RefreshTokenConfig struct {
	Secret     string        `envconfig:"REFRESH_TOKEN_SECRET" required:"true"`
	Expiration time.Duration `envconfig:"REFRESH_TOKEN_EXPIRATION" default:"60m"`
}

func Get() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
