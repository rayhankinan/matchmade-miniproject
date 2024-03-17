package healthcheck

import (
	"net/http"

	"service/internal/handler"

	"github.com/labstack/echo/v4"
)

type IPingHanlder interface {
	PingHandler(c echo.Context) error
}

type PingHandler struct{}

func (h *PingHandler) PingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, handler.SuccessResponse{
		Data: "pong",
	})
}

func NewHandle() (handler IPingHanlder) {
	handler = &PingHandler{}
	return
}
