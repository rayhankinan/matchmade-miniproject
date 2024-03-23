package usecase

import (
	"fmt"
	"service/internal/models"

	"github.com/google/uuid"
)

type WatchlistUseCase struct {
	MovieRepo models.MovieRepository
}

func NewWatchlistUseCase(movieRepo models.MovieRepository) *WatchlistUseCase {
	return &WatchlistUseCase{
		MovieRepo: movieRepo,
	}
}

func (w *WatchlistUseCase) AddMovie(movie models.Movie, userID uuid.UUID) (models.Movie, error) {
	exist, err := w.MovieRepo.IsExist(userID, movie.MovieID)
	if err != nil {
		return models.Movie{}, err
	}

	if exist {
		return models.Movie{}, fmt.Errorf("Movie already exists in watchlist")
	}

	movie.MID = uuid.New()
	movie.UserID = userID

	err = w.MovieRepo.Create(&movie)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func (w *WatchlistUseCase) RemoveMovie(userID uuid.UUID, movieID int64) error {
	return w.MovieRepo.Delete(userID, movieID)
}

func (w *WatchlistUseCase) GetMovies(userID uuid.UUID, title string, page int, pageSize int) (map[string]interface{}, error) {
	movie, err := w.MovieRepo.FindByUserID(userID, title, page, pageSize)
	if err != nil {
		return map[string]interface{}{}, err
	}

	totalResults, err := w.MovieRepo.CountByUserID(userID, title)
	if err != nil {
		return map[string]interface{}{}, err
	}

	totalPages := totalResults / int64(pageSize)
	if totalResults%int64(pageSize) > 0 {
		totalPages++
	}

	return map[string]interface{}{
		"total_results": totalResults,
		"total_pages":   totalPages,
		"movies":        movie,
	}, nil
}

func (w *WatchlistUseCase) GiveRating(userID uuid.UUID, movieID int64, rating int64) error {
	_, err := w.MovieRepo.FindByID(userID, movieID)
	if err != nil {
		return err
	}

	return w.MovieRepo.UpdateRating(userID, movieID, rating)
}

func (w *WatchlistUseCase) IsExist(userID uuid.UUID, movieID int64) (bool, error) {
	return w.MovieRepo.IsExist(userID, movieID)
}
