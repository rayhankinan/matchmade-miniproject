package service

import (
	"service/internal/models"
	"service/internal/utils"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(user models.User) (models.User, error)
	Login(email string, password string) (string, error)
}

type GormAuthService struct {
	DB *gorm.DB
}

func (c *GormAuthService) Register(user models.User) (models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = hashedPassword

	err = c.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *GormAuthService) Login(email string, password string) (string, error) {
	var user models.User
	err := c.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
