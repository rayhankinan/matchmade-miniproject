package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	MID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID `gorm:"not null"`
	Title       string    `gorm:"uniqueIndex;not null"`
	Image       sql.NullString
	ReleaseDate string
	Summary     string
	Genre       string
	Rating      sql.NullInt64
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User User `gorm:"foreignKey:UserID;references:UID"`
}

type MovieRepository interface {
	Create(movie *Movie) error
	Delete(id uuid.UUID) error
	FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]Movie, error)
	FindByID(id uuid.UUID) (*Movie, error)
	UpdateRating(id uuid.UUID, rating int16) error
}
