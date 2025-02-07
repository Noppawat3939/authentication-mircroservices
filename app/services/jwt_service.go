package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken(payload map[string]interface{}, expiredInHours int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	exp := time.Now().Add(time.Duration(expiredInHours) * time.Hour).Unix()

	token.Claims.(jwt.MapClaims)["exp"] = exp

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
