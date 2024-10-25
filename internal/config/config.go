package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token    string         `envconfig:"TOKEN"`
	Postgres PostgresConfig `envconfig:"POSTGRES"`
}

type PostgresConfig struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}

func New() Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return cfg
}
