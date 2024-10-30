package config

import (
	"os"

	"github.com/joho/godotenv"
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

func LoadConfig() (Config, error) {
	env, ok := os.LookupEnv("ENV")
	if ok && env != "" {
		if err := godotenv.Load(); err != nil {
			return Config{}, err
		}
	}

	var cfg Config
	envconfig.MustProcess("", &cfg)

	return cfg, nil
}
