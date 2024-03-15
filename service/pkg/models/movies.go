package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	ImgLink     string
	ReleaseDate string
	Summary     string
	Genre       string
	Rating      *int
	User        User `gorm:"foreignKey:UserID"`
}
