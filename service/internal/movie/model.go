package movie

import (
	"service/internal/user"

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
	User        user.User `gorm:"foreignKey:UserID"`
}
