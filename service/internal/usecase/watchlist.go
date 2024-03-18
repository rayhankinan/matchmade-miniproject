package usecase

import (
	"service/internal/models"

	"github.com/google/uuid"
)

type WatchlistUseCase struct {
	MovieRepo models.MovieRepository
	UserRepo  models.UserRepository
}

func NewWatchlistUseCase(movieRepo models.MovieRepository, userRepo models.UserRepository) *WatchlistUseCase {
	return &WatchlistUseCase{
		MovieRepo: movieRepo,
		UserRepo:  userRepo,
	}
}

func (w *WatchlistUseCase) AddMovie(movie models.Movie, userID uuid.UUID) (models.Movie, error) {
	movie.ID = uuid.New()

	_, err := w.UserRepo.FindById(userID)
	if err != nil {
		return models.Movie{}, err
	}

	movie.UserID = userID

	err = w.MovieRepo.Create(&movie)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func (w *WatchlistUseCase) RemoveMovie(movieID uuid.UUID, userID uuid.UUID) error {
	_, err := w.UserRepo.FindById(userID)
	if err != nil {
		return err
	}

	err = w.MovieRepo.Delete(movieID)
	if err != nil {
		return err
	}

	return nil
}
