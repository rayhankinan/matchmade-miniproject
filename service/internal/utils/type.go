package utils

import (
	"net/http"
	"time"

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
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	cookie.Secure = true

	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1 // Setting MaxAge to -1 forces the browser to delete the cookie immediately.
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
