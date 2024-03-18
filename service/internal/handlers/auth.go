package handlers

import (
	"log"
	"net/http"

	"service/internal/models"
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
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request"})
	}

	user, err = h.AuthUseCase.Register(user)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, utils.ErrorResponse{Message: "Internal server error"})
	}

	log.Println("User registered successfully")

	return utils.SendResponse(c, http.StatusCreated, utils.SuccessResponse{Data: user})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var LoginRequest struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	err := c.Bind(&LoginRequest)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, utils.ErrorResponse{Message: "Invalid request"})
	}

	token, err := h.AuthUseCase.Login(LoginRequest.Identifier, LoginRequest.Password)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, utils.ErrorResponse{Message: "Invalid credentials"})
	}

	log.Println("User logged in successfully")

	utils.SetCookie(c, "AUTH_TOKEN", token)

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse{})
}
