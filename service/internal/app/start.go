package app

import (
	"service/internal/infrastructure"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func NewApp() *App {
	infrastructure.Init()
	db := infrastructure.Create()
	return &App{
		DB: db,
	}
}
