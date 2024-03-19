package usecase

import (
	"service/internal/models"
	"service/internal/repositories/mocks/user"
	"service/internal/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthUseCase_Register(t *testing.T) {
	mockRepo := new(user.MockUserRepo)
	authUseCase := NewAuthUseCase(mockRepo)
	testUser := models.User{
		Username: "testUser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Set expectation
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	// Execute
	user, err := authUseCase.Register(testUser)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user.UID)
	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_Login(t *testing.T) {
	mockRepo := new(user.MockUserRepo)
	authUseCase := NewAuthUseCase(mockRepo)
	testEmail := "test@example.com"
	testPassword := "password"

	hashedPassword, _ := utils.HashPassword(testPassword)
	testUser := models.User{UID: uuid.New(), Email: testEmail, Password: hashedPassword}

	mockRepo.On("FindByEmailOrUsername", testEmail).Return(&testUser, nil)

	token, err := authUseCase.Login(testEmail, testPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
