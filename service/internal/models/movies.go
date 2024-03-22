package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	MID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"not null"`
	MovieID   int64     `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Image     sql.NullString
	Rating    sql.NullInt64
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;references:UID"`
}

type MovieRepository interface {
	Create(movie *Movie) error
	Delete(mid uuid.UUID) error
	FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]Movie, error)
	FindByID(mid uuid.UUID) (*Movie, error)
	UpdateRating(mid uuid.UUID, rating int16) error
	IsExist(userID uuid.UUID, movieID int64) (bool, error)
}
