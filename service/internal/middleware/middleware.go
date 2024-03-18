package middleware

import (
	"log"
	"net/http"
	"strings"

	"service/internal/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		AuthHeader := c.Request().Header.Get("Authorization")

		if AuthHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
		}

		tokenString := strings.TrimPrefix(AuthHeader, "Bearer ")
		if tokenString == AuthHeader {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing Bearer token")
		}

		config, err := config.LoadEnvironment()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error loading environment variables")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.JWTSecret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		log.Println("User authenticated")
		return next(c)
	}
}
