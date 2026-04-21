package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("JWT_SECRET")

func GenerateToken(userID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
