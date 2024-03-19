package middleware

import (
	"log"
	"net/http"

	"service/internal/config"
	"service/internal/utils"

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

		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			utils.DeleteCookie(c, "AUTH_TOKEN") // Delete the cookie if the token is invalid
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			utils.DeleteCookie(c, "AUTH_TOKEN") // Delete the cookie if the token is invalid
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("userID", claims["userID"])
		c.Set("email", claims["email"])
		c.Set("username", claims["username"])

		log.Println("User", claims["userID"], "authenticated")

		return next(c)
	}
}
