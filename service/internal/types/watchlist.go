package types

import (
	"database/sql"
	"service/internal/models"

	"github.com/google/uuid"
)

type MovieRequest struct {
	MovieID     string  `json:"movieID"`
	Title       string  `json:"title"`
	Image       *string `json:"image,omitempty"`
	ReleaseDate string  `json:"releaseDate"`
	Summary     string  `json:"summary"`
	Genre       string  `json:"genre"`
	Rating      *int16  `json:"rating,omitempty"`
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
		UserID:      userID,
		MovieID:     m.MovieID,
		Title:       m.Title,
		Image:       image,
		ReleaseDate: m.ReleaseDate,
		Summary:     m.Summary,
		Genre:       m.Genre,
		Rating:      rating, // Ensure your Movie struct in the domain model expects sql.NullInt64 for Rating
	}, nil
}
