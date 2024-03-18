package middleware

import (
	"log"
	"net/http"

	"service/internal/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("AUTH_TOKEN")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenString := token.Value

		config, err := config.LoadEnvironment()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error loading environment variables")
		}

		parsedToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.JWTSecret), nil
		})

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			c.Set("userID", claims["email"])
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		log.Println("User authenticated")
		return next(c)
	}
}
