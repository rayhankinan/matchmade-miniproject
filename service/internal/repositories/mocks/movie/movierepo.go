package movie

import (
	"service/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// Mock MovieRepository
type MockMovieRepo struct {
	mock.Mock
}

func (_m *MockMovieRepo) Create(movie *models.Movie) error {
	ret := _m.Called(movie)
	return ret.Error(0)
}

func (_m *MockMovieRepo) Delete(id int64) error {
	args := _m.Called(id)
	return args.Error(0)
}

func (_m *MockMovieRepo) FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	ret := _m.Called(userID, title, page, pageSize)
	return ret.Get(0).([]models.Movie), ret.Error(1)
}

func (_m *MockMovieRepo) FindByID(id int64) (*models.Movie, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*models.Movie), ret.Error(1)
}

func (_m *MockMovieRepo) UpdateRating(id int64, rating int16) error {
	args := _m.Called(id, rating)
	return args.Error(0)
}

func (_m *MockMovieRepo) IsExist(userID uuid.UUID, movieID int64) (bool, error) {
	ret := _m.Called(userID, movieID)
	return ret.Bool(0), ret.Error(1)
}
