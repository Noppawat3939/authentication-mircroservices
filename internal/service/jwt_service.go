package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

func NewJwtService(secretKey string) *JWTService {
	return &JWTService{SecretKey: secretKey}
}

func (service *JWTService) GenerateToken(claims map[string]interface{}, expiredHour *int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	if expiredHour != nil {
		expirationTime := time.Now().Add(time.Duration(*expiredHour) * time.Hour)
		token.Claims.(jwt.MapClaims)["exp"] = expirationTime.Unix()
	}

	fmt.Print(claims)

	return token.SignedString([]byte(service.SecretKey))
}
