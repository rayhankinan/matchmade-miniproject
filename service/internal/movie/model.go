package movie

import (
	"database/sql"
	"service/internal/user"

	"gorm.io/gorm"
)

// TODO: Add necessary fields to preview the movie
type Movie struct {
	gorm.Model

	UserID  uint `gorm:"not null"`
	MovieID uint `gorm:"not null"`
	Rating  sql.Null[uint16]

	User user.User `gorm:"foreignKey:UserID;references:ID"`
}
