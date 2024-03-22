package types

import (
	"database/sql"
	"service/internal/models"

	"github.com/google/uuid"
)

type MovieRequest struct {
	MovieID int64   `json:"movieID" validate:"required"`
	Title   string  `json:"title" validate:"required"`
	Image   *string `json:"image,omitempty"`
	Rating  *int16  `json:"rating,omitempty"`
}

type RatingRequest struct {
	Rating int16 `json:"rating" validate:"required,gte=1,lte=5"`
}

func (m *MovieRequest) ToMovie(userID uuid.UUID) (models.Movie, error) {
	var image sql.NullString
	if m.Image != nil {
		image = sql.NullString{String: *m.Image, Valid: true}
	} else {
		image = sql.NullString{Valid: false}
	}

	var rating sql.NullInt64
	if m.Rating != nil {
		rating = sql.NullInt64{Int64: int64(*m.Rating), Valid: true}
	} else {
		rating = sql.NullInt64{Valid: false}
	}

	return models.Movie{
		UserID:  userID,
		MovieID: m.MovieID,
		Title:   m.Title,
		Image:   image,
		Rating:  rating,
	}, nil
}
