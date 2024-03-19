package handlers

import (
	"log"
	"net/http"
	"strconv"

	"service/internal/types"
	"service/internal/usecase"
	"service/internal/utils"

	"github.com/go-playground/validator/v10"
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
	var req types.MovieRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movie, _ := req.ToMovie(UserID)

	movie, err = h.WatchlistUseCase.AddMovie(movie, UserID)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Movie already exists in watchlist"})
	}

	log.Println("Movie added to watchlist successfully")

	return utils.SendResponse(c, http.StatusCreated, types.SuccessResponse{Data: movie})
}

func (h *WatchlistHandler) RemoveMovieFromWatchlist(c echo.Context) error {
	movieID := c.Param("id")
	MovieID, err := uuid.Parse(movieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	err = h.WatchlistUseCase.RemoveMovie(MovieID, UserID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "The movie does not exist in watchlist"})
	}

	log.Println("Movie removed from watchlist successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: "Movie removed from watchlist"})
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
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movies, err := h.WatchlistUseCase.GetMovies(UserID, title, page, pageSize)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to get movies from watchlist"})
	}

	log.Println("Movies retrieved successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: movies})
}

func (h *WatchlistHandler) GetMovieDetail(c echo.Context) error {
	movieID := c.Param("id")
	MovieID, err := uuid.Parse(movieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movie, err := h.WatchlistUseCase.GetMovieDetail(MovieID, UserID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "The movie does not exist in watchlist"})
	}

	log.Println("Movie detail retrieved successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: movie})
}

func (h *WatchlistHandler) GiveRating(c echo.Context) error {
	movieID := c.Param("id")
	MovieID, err := uuid.Parse(movieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	var req types.RatingRequest

	err = c.Bind(&req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	err = h.WatchlistUseCase.GiveRating(MovieID, UserID, req.Rating)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to give rating to movie"})
	}

	log.Println("Rating given to movie successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: "Rating given to movie successfully"})
}
