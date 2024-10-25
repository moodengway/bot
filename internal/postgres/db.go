package postgres

import (
	"fmt"

	"github.com/nonya123456/connect4/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
