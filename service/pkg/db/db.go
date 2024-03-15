package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"service/pkg/config"
	"service/pkg/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Config.Host, config.Config.User, config.Config.Password, config.Config.Name, config.Config.Port, config.Config.SSLMode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	err = DB.AutoMigrate(&models.User{}, &models.Movie{})
	if err != nil {
		log.Fatal(err)
	}
}
