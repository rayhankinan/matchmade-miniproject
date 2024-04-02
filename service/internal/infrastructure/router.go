package infrastructure

import (
	"service/internal/handlers"
	"service/internal/middleware"
	"service/internal/repositories"
	"service/internal/usecase"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func CreateRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	userRepo := repositories.NewGormUserRepo(db)
	authUsecase := usecase.NewAuthUseCase(userRepo)
	authHandler := handlers.NewAuthHandler(authUsecase)

	movieRepo := repositories.NewGormMovieRepo(db)
	watchlistUsecase := usecase.NewWatchlistUseCase(movieRepo)
	watchlistHandler := handlers.NewWatchlistHandler(watchlistUsecase)

	userRoute := e.Group("/users")
	userRoute.POST("/register", authHandler.Register)
	userRoute.POST("/login", authHandler.Login)
	userRoute.POST("/logout", authHandler.Logout)

	movieRoute := e.Group("/movies")
	movieRoute.GET("/watchlist/exist/:id", watchlistHandler.IsMovieInWatchlist, middleware.JWTAuthMiddleware)
	movieRoute.GET("/watchlist", watchlistHandler.GetMovies, middleware.JWTAuthMiddleware)
	movieRoute.POST("/add", watchlistHandler.AddMovieToWatchlist, middleware.JWTAuthMiddleware)
	movieRoute.GET("/rate/:id", watchlistHandler.GetRating, middleware.JWTAuthMiddleware)
	movieRoute.PATCH("/rate/:id", watchlistHandler.GiveRating, middleware.JWTAuthMiddleware)
	movieRoute.DELETE("/remove/:id", watchlistHandler.RemoveMovieFromWatchlist, middleware.JWTAuthMiddleware)
	movieRoute.GET("/tags/:id", watchlistHandler.GetTags, middleware.JWTAuthMiddleware)

	return e
}
