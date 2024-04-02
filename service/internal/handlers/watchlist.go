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

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
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
	MovieID, err := strconv.ParseInt(movieID, 10, 64)
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

	err = h.WatchlistUseCase.RemoveMovie(UserID, MovieID)
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
		page = 1
	}

	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = 10
	}

	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	result, err := h.WatchlistUseCase.GetMovies(UserID, title, page, pageSize)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to get movies from watchlist"})
	}

	log.Println("Movies retrieved successfully")

	return utils.SendPaginatedResponse(c, http.StatusOK, types.PaginatedResponse{
		Data:         result["movies"],
		Page:         int64(page),
		TotalPages:   result["total_pages"].(int64),
		TotalResults: result["total_results"].(int64),
	})
}

func (h *WatchlistHandler) GiveRating(c echo.Context) error {
	movieID := c.Param("id")
	MovieID, err := strconv.ParseInt(movieID, 10, 64)
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

	err = h.WatchlistUseCase.GiveRating(UserID, MovieID, req.Rating)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to give rating to movie"})
	}

	log.Println("Rating given to movie successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: "Rating given to movie successfully"})
}

func (h *WatchlistHandler) IsMovieInWatchlist(c echo.Context) error {
	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movieID := c.Param("id")
	MovieID, err := strconv.ParseInt(movieID, 10, 64)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	exist, err := h.WatchlistUseCase.IsExist(UserID, MovieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to check if movie is in watchlist"})
	}

	log.Println("Movie existence in watchlist checked successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: exist})
}

func (h *WatchlistHandler) GetRating(c echo.Context) error {
	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movieID := c.Param("id")
	MovieID, err := strconv.ParseInt(movieID, 10, 64)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	rating, err := h.WatchlistUseCase.GetRating(UserID, MovieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to get rating of movie"})
	}

	log.Println("Movie rating retrieved successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: rating})
}

func (h *WatchlistHandler) GetTags(c echo.Context) error {
	userID := c.Get("userID").(string)
	UserID, err := uuid.Parse(userID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Not authorized to perform this action"})
	}

	movieID := c.Param("id")
	MovieID, err := strconv.ParseInt(movieID, 10, 64)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid movie ID"})
	}

	tags, err := h.WatchlistUseCase.GetTags(UserID, MovieID)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Failed to get tags of movie"})
	}

	log.Println("Movie tags retrieved successfully")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{Data: tags})
}
