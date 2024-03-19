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
	movie.MID = uuid.New()

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

func (w *WatchlistUseCase) GetMovies(userID uuid.UUID, title string, page int, pageSize int) ([]models.Movie, error) {
	_, err := w.UserRepo.FindById(userID)
	if err != nil {
		return nil, err
	}

	return w.MovieRepo.FindByUserID(userID, title, page, pageSize)
}

func (w *WatchlistUseCase) GetMovieDetail(movieID uuid.UUID, userID uuid.UUID) (*models.Movie, error) {
	_, err := w.UserRepo.FindById(userID)
	if err != nil {
		return nil, err
	}

	return w.MovieRepo.FindByID(movieID)
}

func (w *WatchlistUseCase) GiveRating(movieID uuid.UUID, userID uuid.UUID, rating int16) error {
	_, err := w.UserRepo.FindById(userID)
	if err != nil {
		return err
	}

	return w.MovieRepo.UpdateRating(movieID, rating)
}
