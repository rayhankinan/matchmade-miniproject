package repositories

import (
	"database/sql"
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

func (c *GormMovieRepo) Delete(mid uuid.UUID) error {
	return c.DB.Delete(&models.Movie{}, mid).Error
}

func (c *GormMovieRepo) FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	var movies []models.Movie
	query := c.DB.Where("user_id = ?", userID)

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	offset := (page - 1) * pageSize
	return movies, query.Offset(offset).Limit(pageSize).Find(&movies).Error
}

func (c *GormMovieRepo) FindByID(mid uuid.UUID) (*models.Movie, error) {
	var movie models.Movie
	err := c.DB.First(&movie, "m_id = ?", mid).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (c *GormMovieRepo) UpdateRating(mid uuid.UUID, rating int16) error {
	result := c.DB.Model(&models.Movie{}).Where("m_id = ?", mid).Updates(
		map[string]interface{}{
			"rating":     sql.NullInt64{Int64: int64(rating), Valid: true},
			"updated_at": time.Now(),
		})
	return result.Error
}

func (c *GormMovieRepo) IsExist(userID uuid.UUID, movieID string) (bool, error) {
	var count int64
	err := c.DB.Model(&models.Movie{}).Where("user_id = ? AND movie_id = ?", userID, movieID).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
