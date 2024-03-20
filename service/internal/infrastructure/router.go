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
	userUsecase := usecase.NewAuthUseCase(userRepo)
	userHandler := handlers.NewAuthHandler(userUsecase)

	movieRepo := repositories.NewGormMovieRepo(db)
	movieUsecase := usecase.NewWatchlistUseCase(movieRepo, userRepo)
	movieHandler := handlers.NewWatchlistHandler(movieUsecase)

	userRoute := e.Group("/users")
	userRoute.POST("/register", userHandler.Register)
	userRoute.POST("/login", userHandler.Login)
	userRoute.POST("/logout", userHandler.Logout)

	movieRoute := e.Group("/movies", middleware.JWTAuthMiddleware)
	movieRoute.POST("/add", movieHandler.AddMovieToWatchlist, middleware.JWTAuthMiddleware)
	movieRoute.DELETE("/remove/:id", movieHandler.RemoveMovieFromWatchlist, middleware.JWTAuthMiddleware)
	movieRoute.GET("/watchlist", movieHandler.GetMovies, middleware.JWTAuthMiddleware)
	movieRoute.GET("/watchlist/:id", movieHandler.GetMovieDetail, middleware.JWTAuthMiddleware)
	movieRoute.PATCH("/watchlist/rate/:id", movieHandler.GiveRating, middleware.JWTAuthMiddleware)
	movieRoute.GET("/watchlist/exist/:id", movieHandler.IsMovieInWatchlist, middleware.JWTAuthMiddleware)

	return e
}
