package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func SetCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Domain = ".localhost"
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true

	c.SetCookie(cookie)
}

func SendResponse(c echo.Context, code int, data SuccessResponse) error {
	return c.JSON(code, data)
}

func SendError(c echo.Context, code int, message ErrorResponse) error {
	return c.JSON(code, message)
}
