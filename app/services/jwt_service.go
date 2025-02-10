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

func GenerateRefreshToken(payload map[string]interface{}) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	exp := time.Now().Add(time.Duration(72) * time.Hour).Unix()

	token.Claims.(jwt.MapClaims)["exp"] = exp

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil, err
	}

	return token.Valid, claims, nil
}
