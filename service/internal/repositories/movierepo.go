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

func (c *GormMovieRepo) Delete(userID uuid.UUID, movieID int64) error {
	return c.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Movie{}).Error
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

func (c *GormMovieRepo) FindByID(userID uuid.UUID, movieID int64) (*models.Movie, error) {
	var movie models.Movie
	err := c.DB.First(&movie, "user_id = ? AND movie_id = ?", userID, movieID).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (c *GormMovieRepo) UpdateRating(userID uuid.UUID, movieID int64, rating int64) error {
	result := c.DB.Model(&models.Movie{}).Where("user_id = ? AND movie_id = ?", userID, movieID).Updates(
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

func (c *GormMovieRepo) CountByUserID(userID uuid.UUID, title string) (int64, error) {
	var count int64
	query := c.DB.Model(&models.Movie{}).Where("user_id = ?", userID)

	if title != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", title))
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *GormMovieRepo) GetRating(userID uuid.UUID, movieID int64) (int64, error) {
	var movie models.Movie
	err := c.DB.Select("rating").First(&movie, "user_id = ? AND movie_id = ?", userID, movieID).Error
	if err != nil {
		return 0, err
	}

	return movie.Rating.Int64, nil
}

func (c *GormMovieRepo) GetTags(userID uuid.UUID, movieID int64) ([]string, error) {
	var movie models.Movie
	err := c.DB.Select("tags").First(&movie, "user_id = ? AND movie_id = ?", userID, movieID).Error
	if err != nil {
		return nil, err
	}

	return movie.Tags, nil
}
