package healthcheck

import (
	"net/http"

	"service/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IHealthcheckHandler interface {
	PingHandler(c echo.Context) error
}

type HealthcheckHandler struct {
	usecase IHealthcheckUseCase
	logger  *logrus.Logger
}

// Handle the request and return the response
func (h *HealthcheckHandler) PingHandler(c echo.Context) (err error) {
	// Create a new context
	ctx := c.Request().Context()

	// Ping the database
	err = h.usecase.Ping(ctx)
	if err != nil {
		// Log the error
		h.logger.Errorf("Failed to ping the database: %v", err)

		// Return the response
		err = c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "Failed to ping the database"})
		return
	}

	// Log the success
	h.logger.Info("Successfully pinged the database")

	// Return the response
	err = c.JSON(http.StatusOK, handler.SuccessResponse{Data: "Pong"})
	return
}

func NewHandler(usecase IHealthcheckUseCase, logger *logrus.Logger) (handler IHealthcheckHandler) {
	handler = &HealthcheckHandler{
		usecase,
		logger,
	}
	return
}
