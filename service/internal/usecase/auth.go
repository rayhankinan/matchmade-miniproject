package usecase

import (
	"service/internal/models"
	"service/internal/utils"

	"github.com/google/uuid"
)

type AuthUseCase struct {
	UserRepo models.UserRepository
}

func NewAuthUseCase(userRepo models.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		UserRepo: userRepo,
	}
}

func (a *AuthUseCase) Register(user models.User) (models.User, error) {
	user.UID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = hashedPassword

	err = a.UserRepo.Create(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a *AuthUseCase) Login(identifier string, password string) (string, error) {
	user, err := a.UserRepo.FindByEmailOrUsername(identifier)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.UID, user.Email, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
