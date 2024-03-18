package infrastructure

import (
	"net/http"

	"service/internal/handlers"
	"service/internal/middleware"
	"service/internal/repositories"
	"service/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	userRepo := repositories.NewGormUserRepo(db)
	userUsecase := usecase.NewAuthUseCase(userRepo)
	userHandler := handlers.NewAuthHandler(userUsecase)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	secure := e.Group("/v1")
	secure.Use(middleware.JWTAuthMiddleware)
	secure.GET("/tes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})

	return e
}
