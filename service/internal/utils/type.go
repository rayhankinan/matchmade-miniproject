package utils

import (
	"database/sql"
	"net/http"
	"service/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MovieRequest struct {
	Title       string  `json:"title"`
	Image       *string `json:"image,omitempty"`
	ReleaseDate string  `json:"releaseDate"`
	Summary     string  `json:"summary"`
	Genre       string  `json:"genre"`
	Rating      *int16  `json:"rating,omitempty"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
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
		Title:       m.Title,
		Image:       image,
		ReleaseDate: m.ReleaseDate,
		Summary:     m.Summary,
		Genre:       m.Genre,
		Rating:      rating, // Ensure your Movie struct in the domain model expects sql.NullInt64 for Rating
	}, nil
}

func SetCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	cookie.Secure = true

	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1 // Setting MaxAge to -1 forces the browser to delete the cookie immediately.
	cookie.HttpOnly = true
	cookie.Secure = true

	c.SetCookie(cookie)

}

func SendResponse(c echo.Context, code int, data SuccessResponse) error {
	return c.JSON(code, data)
}

func SendError(c echo.Context, code int, message ErrorResponse) error {
	return c.JSON(code, message)
}
