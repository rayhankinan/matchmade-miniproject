package infrastructure

import (
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
	movieRoute.PATCH("/watchlist/rate/:id", watchlistHandler.GiveRating, middleware.JWTAuthMiddleware)
	movieRoute.DELETE("/remove/:id", watchlistHandler.RemoveMovieFromWatchlist, middleware.JWTAuthMiddleware)

	return e
}
