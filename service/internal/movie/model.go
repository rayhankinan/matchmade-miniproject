package movie

import (
	"database/sql"
	"service/internal/user"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model

	UserID  uint   `gorm:"not null"`
	MovieID uint   `gorm:"not null"`
	Title   string `gorm:"not null"`
	Image   sql.Null[string]
	Rating  sql.Null[uint16]

	User user.User `gorm:"foreignKey:UserID;references:ID"`
}
