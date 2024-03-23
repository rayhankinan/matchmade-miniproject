package usecase

import (
	"service/internal/models"
	"service/internal/repositories/mocks/movie"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddMovie(t *testing.T) {
	mockMovieRepo := new(movie.MockMovieRepo)
	watchlistUseCase := NewWatchlistUseCase(mockMovieRepo)

	userID := uuid.New()
	now := time.Now()

	testMovie := models.Movie{
		UserID:    userID,
		MovieID:   123,
		Title:     "Test Movie",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Setup expectations
	mockMovieRepo.On("IsExist", userID, testMovie.MovieID).Return(false, nil)
	mockMovieRepo.On("Create", mock.AnythingOfType("*models.Movie")).Return(nil)

	// Execute the method under test
	result, err := watchlistUseCase.AddMovie(testMovie, userID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, testMovie.Title, result.Title)
	mockMovieRepo.AssertExpectations(t)
}

func TestRemoveMovie(t *testing.T) {
	mockMovieRepo := new(movie.MockMovieRepo)
	watchlistUseCase := NewWatchlistUseCase(mockMovieRepo)
	userID := uuid.New()
	movieID := int64(123)

	// Setup expectations
	mockMovieRepo.On("Delete", userID, movieID).Return(nil)

	// Execute the method under test
	err := watchlistUseCase.RemoveMovie(userID, movieID)

	// Assertions
	assert.NoError(t, err)
	mockMovieRepo.AssertExpectations(t)
}

func TestGetMovies(t *testing.T) {
	mockMovieRepo := new(movie.MockMovieRepo)
	watchlistUseCase := NewWatchlistUseCase(mockMovieRepo)
	userID := uuid.New()
	now := time.Now()

	testMovie := models.Movie{
		UserID:    userID,
		MovieID:   123,
		Title:     "Test Movie",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Setup expectations
	mockMovieRepo.On("FindByUserID", userID, testMovie.Title, 1, 10).Return([]models.Movie{testMovie}, nil)

	// Execute the method under test
	result, err := watchlistUseCase.GetMovies(userID, testMovie.Title, 1, 10)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, testMovie.Title, result[0].Title)
	mockMovieRepo.AssertExpectations(t)
}

func TestGiveRating(t *testing.T) {
	mockMovieRepo := new(movie.MockMovieRepo)
	watchlistUseCase := NewWatchlistUseCase(mockMovieRepo)
	userID := uuid.New()
	movieID := int64(123)
	rating := int64(5)

	// Setup expectations
	mockMovieRepo.On("FindByID", userID, movieID).Return(&models.Movie{}, nil)
	mockMovieRepo.On("UpdateRating", userID, movieID, rating).Return(nil)

	// Execute the method under test
	err := watchlistUseCase.GiveRating(userID, movieID, rating)

	// Assertions
	assert.NoError(t, err)
	mockMovieRepo.AssertExpectations(t)
}

func TestIsExist(t *testing.T) {
	mockMovieRepo := new(movie.MockMovieRepo)
	watchlistUseCase := NewWatchlistUseCase(mockMovieRepo)
	userID := uuid.New()
	movieID := int64(123)

	// Setup expectations
	mockMovieRepo.On("IsExist", userID, movieID).Return(true, nil)

	// Execute the method under test
	result, err := watchlistUseCase.IsExist(userID, movieID)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, result)
	mockMovieRepo.AssertExpectations(t)
}
