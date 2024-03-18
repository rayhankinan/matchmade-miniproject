package repositories

import (
	"service/internal/models"

	"github.com/google/uuid"
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

func (c *GormUserRepo) FindByEmailOrUsername(identifier string) (*models.User, error) {
	var user models.User
	err := c.DB.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *GormUserRepo) FindById(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := c.DB.Find(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
