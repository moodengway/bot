package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token string `envconfig:"TOKEN"`
}

func New() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
