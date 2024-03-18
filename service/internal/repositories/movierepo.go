package repositories

import (
	"service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormMovieRepo struct {
	DB *gorm.DB
}

func NewGormMovieRepo(db *gorm.DB) models.MovieRepository {
	return &GormMovieRepo{
		DB: db,
	}
}

func (c *GormMovieRepo) Create(movie *models.Movie) error {
	return c.DB.Create(&movie).Error
}

func (c *GormMovieRepo) Delete(id uuid.UUID) error {
	return c.DB.Delete(&models.Movie{}, id).Error
}
