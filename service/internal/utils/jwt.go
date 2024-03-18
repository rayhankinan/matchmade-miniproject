package utils

import (
	"time"

	"service/internal/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userId uuid.UUID) (string, error) {
	config, err := config.LoadEnvironment()
	if err != nil {
		return "", err
	}

	jwtKey := []byte(config.JWTSecret)

	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
