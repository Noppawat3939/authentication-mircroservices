package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

func (service *JWTService) GenerateToken(email, username string, expiredHour *int64) (string, error) {
	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
	}

	if expiredHour != nil {
		claims["exp"] = time.Now().Add(time.Duration(*expiredHour) * time.Hour).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.SecretKey))
}
