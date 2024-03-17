package app

import (
	"fmt"

	"service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn(cfg config.EnvironmentConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName),
		),
		&gorm.Config{},
	)
	return
}
