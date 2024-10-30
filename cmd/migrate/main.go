package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nonya123456/connect4/internal/config"
	internalpostgres "github.com/nonya123456/connect4/internal/postgres"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Panic("error loading config", zap.Error(err))
	}

	dsn := internalpostgres.ToDSN(cfg.Postgres)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Panic("error openning sql database", zap.Error(err))
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Panic("error creating driver", zap.Error(err))
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		logger.Panic("error creating migration", zap.Error(err))
	}

	_ = m.Up()
}
