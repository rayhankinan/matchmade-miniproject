package handlers

import (
	"log"
	"net/http"

	"service/internal/models"
	"service/internal/types"
	"service/internal/usecase"
	"service/internal/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseAuthUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUseAuthUseCase,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req types.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	user := models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	res, err := h.AuthUseCase.Register(user)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusInternalServerError, types.ErrorResponse{Message: "Internal server error: " + err.Error()})
	}

	log.Println("User registered successfully")

	return utils.SendResponse(c, http.StatusCreated, types.SuccessResponse{Data: res})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req types.LoginRequest

	err := c.Bind(&req)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusBadRequest, types.ErrorResponse{Message: "Invalid request: Please provide valid data"})
	}

	token, err := h.AuthUseCase.Login(req.Identifier, req.Password)
	if err != nil {
		log.Println(err)
		return utils.SendError(c, http.StatusUnauthorized, types.ErrorResponse{Message: "Invalid credentials"})
	}

	log.Println("User logged in successfully")

	utils.SetCookie(c, "AUTH_TOKEN", token)

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	utils.DeleteCookie(c, "AUTH_TOKEN")

	return utils.SendResponse(c, http.StatusOK, types.SuccessResponse{})
}
