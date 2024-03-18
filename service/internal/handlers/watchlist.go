package handlers

import (
	"log"
	"net/http"

	"service/internal/usecase"
	"service/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type WatchlistHandler struct {
	WatchlistUseCase *usecase.WatchlistUseCase
}

func NewWatchlistHandler(watchlistUseCase *usecase.WatchlistUseCase) *WatchlistHandler {
	return &WatchlistHandler{
		WatchlistUseCase: watchlistUseCase,
	}
}

func (h *WatchlistHandler) AddMovieToWatchlist(c echo.Context) error {
	var MovieRequest utils.MovieRequest
	err := c.Bind(&MovieRequest)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, utils.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movie, _ := MovieRequest.ToMovie(UserID)

	movie, err = h.WatchlistUseCase.AddMovie(movie, UserID)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to add movie to watchlist"})
	}

	log.Println("Movie added to watchlist successfully")

	return utils.SendResponse(c, http.StatusCreated, utils.SuccessResponse{Data: movie})
}

func (h *WatchlistHandler) RemoveMovieFromWatchlist(c echo.Context) error {
	movieID := c.Param("id")
	MovieID, err := uuid.Parse(movieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid movie ID"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, utils.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	err = h.WatchlistUseCase.RemoveMovie(MovieID, UserID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to remove movie from watchlist"})
	}

	log.Println("Movie removed from watchlist successfully")

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse{Data: "Movie removed from watchlist"})
}
