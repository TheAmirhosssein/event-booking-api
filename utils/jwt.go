package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"user_id":  userId,
		"exp":      time.Now().Add(time.Hour * (24 * 365)).Unix(),
	})

	secretKey := os.Getenv("SECRET_KEY")
	return token.SignedString([]byte(secretKey))
}
