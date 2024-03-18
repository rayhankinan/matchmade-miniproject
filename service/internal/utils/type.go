package utils

import "github.com/labstack/echo/v4"

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendResponse(c echo.Context, code int, data SuccessResponse) error {
	return c.JSON(code, data)
}

func SendError(c echo.Context, code int, message ErrorResponse) error {
	return c.JSON(code, message)
}
