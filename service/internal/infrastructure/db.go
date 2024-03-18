package infrastructure

import (
	"fmt"
	"service/internal/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	config, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DatabaseHost, strconv.Itoa(config.DatabasePort), config.DatabaseUser, config.DatabaseName, config.DatabasePassword)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Create() *gorm.DB {
	return DB
}
