package app

import (
	"service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConn(cfg config.EnvironmentConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	return
}
