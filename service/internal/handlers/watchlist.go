package handlers

import (
	"log"
	"net/http"
	"strconv"

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

func (h *WatchlistHandler) GetMovies(c echo.Context) error {
	title := c.QueryParam("title")
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("pageSize")

	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page = 1 // Default to first page
	}

	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = 10 // Default page size
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, utils.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movies, err := h.WatchlistUseCase.GetMovies(UserID, title, page, pageSize)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to get movies from watchlist"})
	}

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse{Data: movies})
}

func (h *WatchlistHandler) GetMovieDetail(c echo.Context) error {
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

	movie, err := h.WatchlistUseCase.GetMovieDetail(MovieID, UserID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to get movie from watchlist"})
	}

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse{Data: movie})
}

func (h *WatchlistHandler) GiveRating(c echo.Context) error {
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

	var RatingReq struct {
		Rating int16 `json:"rating"`
	}

	err = c.Bind(&RatingReq)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	err = h.WatchlistUseCase.GiveRating(MovieID, UserID, RatingReq.Rating)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Failed to give rating to movie"})
	}

	log.Println("Rating given to movie successfully")

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse{Data: "Rating given to movie successfully"})
}