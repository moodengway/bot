package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
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

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}

func LoadConfig() (AppConfig, error) {
	env, ok := os.LookupEnv("ENV")
	if ok && env != "" {
		if err := godotenv.Load(); err != nil {
			return AppConfig{}, err
		}
	}

	var cfg AppConfig
	envconfig.MustProcess("APP", &cfg)

	return cfg, nil
}
