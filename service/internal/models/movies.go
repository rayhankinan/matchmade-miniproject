package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Image       sql.Null[string]
	ReleaseDate string
	Summary     string
	Genre       string
	Rating      sql.Null[uint16]

	User User `gorm:"foreignKey:UserID;references:ID"`
}
