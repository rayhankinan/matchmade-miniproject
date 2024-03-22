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

func (w *WatchlistUseCase) RemoveMovie(movieID int64) error {
	return w.MovieRepo.Delete(movieID)
}

func (w *WatchlistUseCase) GetMovies(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	return w.MovieRepo.FindByUserID(userID, title, page, pageSize)
}

func (w *WatchlistUseCase) GiveRating(movieID int64, rating int64) error {
	_, err := w.MovieRepo.FindByID(movieID)
	if err != nil {
		return err
	}

	return w.MovieRepo.UpdateRating(movieID, rating)
}

func (w *WatchlistUseCase) IsExist(userID uuid.UUID, movieID int64) (bool, error) {
	return w.MovieRepo.IsExist(userID, movieID)
}
