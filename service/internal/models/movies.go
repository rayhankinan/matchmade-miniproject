package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Image       sql.Null[string]
	ReleaseDate string
	Summary     string
	Genre       string
	Rating      sql.Null[uint16]
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User User `gorm:"foreignKey:UserID;references:ID"`
}

type MovieRepository interface {
	Create(movie *Movie) error
	Delete(id uuid.UUID) error
}
