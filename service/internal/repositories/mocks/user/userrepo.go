package user

import (
	"service/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// Mock UserRepository
type MockUserRepo struct {
	mock.Mock
}

func (_m *MockUserRepo) Create(user *models.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

func (_m *MockUserRepo) FindByEmailOrUsername(identifier string) (*models.User, error) {
	ret := _m.Called(identifier)
	return ret.Get(0).(*models.User), ret.Error(1)
}

func (_m *MockUserRepo) FindById(id uuid.UUID) (*models.User, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*models.User), ret.Error(1)
}
