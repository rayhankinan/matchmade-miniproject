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

func (_m *MockMovieRepo) Delete(userID uuid.UUID, movieID int64) error {
	ret := _m.Called(userID, movieID)
	return ret.Error(0)
}

func (_m *MockMovieRepo) FindByUserID(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	ret := _m.Called(userID, title, page, pageSize)
	return ret.Get(0).([]models.Movie), ret.Error(1)
}

func (_m *MockMovieRepo) FindByID(userID uuid.UUID, movieID int64) (*models.Movie, error) {
	ret := _m.Called(userID, movieID)
	return ret.Get(0).(*models.Movie), ret.Error(1)
}

func (_m *MockMovieRepo) UpdateRating(userID uuid.UUID, movieID int64, rating int64) error {
	ret := _m.Called(userID, movieID, rating)
	return ret.Error(0)
}

func (_m *MockMovieRepo) IsExist(userID uuid.UUID, movieID int64) (bool, error) {
	ret := _m.Called(userID, movieID)
	return ret.Bool(0), ret.Error(1)
}

func (_m *MockMovieRepo) CountByUserID(userID uuid.UUID, title string) (int64, error) {
	ret := _m.Called(userID, title)
	return ret.Get(0).(int64), ret.Error(1)
}

func (_m *MockMovieRepo) GetRating(userID uuid.UUID, movieID int64) (int64, error) {
	ret := _m.Called(userID, movieID)
	return ret.Get(0).(int64), ret.Error(1)
}

func (_m *MockMovieRepo) GetTags(userID uuid.UUID, movieID int64) ([]string, error) {
	ret := _m.Called(userID, movieID)
	return ret.Get(0).([]string), ret.Error(1)
}
