package repositories

import (
	"service/internal/models"

	"gorm.io/gorm"
)

type GormUserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) models.UserRepository {
	return &GormUserRepo{
		DB: db,
	}
}

func (c *GormUserRepo) Create(user *models.User) error {
	return c.DB.Create(&user).Error
}

func (c *GormUserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := c.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
