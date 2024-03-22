package repositories

import (
	"database/sql"
	"fmt"
	"service/internal/models"
	"time"

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

func (c *GormMovieRepo) Delete(mid int64) error {
	return c.DB.Where("movie_id = ?", mid).Delete(&models.Movie{}).Error
}

func (c *GormMovieRepo) FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	var movies []models.Movie
	query := c.DB.Where("user_id = ?", userID)

	if title != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", title))
	}

	offset := (page - 1) * pageSize
	return movies, query.Offset(offset).Limit(pageSize).Find(&movies).Error
}

func (c *GormMovieRepo) FindByID(mid int64) (*models.Movie, error) {
	var movie models.Movie
	err := c.DB.First(&movie, "movie_id = ?", mid).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (c *GormMovieRepo) UpdateRating(mid int64, rating int64) error {
	result := c.DB.Model(&models.Movie{}).Where("movie_id = ?", mid).Updates(
		map[string]interface{}{
			"rating":     sql.NullInt64{Int64: rating, Valid: true},
			"updated_at": time.Now(),
		})
	return result.Error
}

func (c *GormMovieRepo) IsExist(userID uuid.UUID, movieID int64) (bool, error) {
	var count int64
	err := c.DB.Model(&models.Movie{}).Where("user_id = ? AND movie_id = ?", userID, movieID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
