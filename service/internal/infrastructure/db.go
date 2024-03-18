package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"service/internal/config"
)

var DB *gorm.DB

func Init() {
	config, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Create() *gorm.DB {
	return DB
}
