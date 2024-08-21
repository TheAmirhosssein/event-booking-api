package utils

import (
	"errors"
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

func ValidateToken(token string) (map[string]any, error) {
	paredToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid type")
		}
		secretKey := os.Getenv("SECRET_KEY")
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if !paredToken.Valid {
		return nil, errors.New("invalid token")
	}
	return paredToken.Claims.(jwt.MapClaims), nil
}
