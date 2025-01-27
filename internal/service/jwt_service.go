package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey        string
	RefreshSecretKey string
}

func NewJwtService(secretKey string, refreshSecretKey string) *JWTService {
	return &JWTService{SecretKey: secretKey, RefreshSecretKey: refreshSecretKey}
}

func (service *JWTService) GenerateAccessToken(claims map[string]interface{}, expiredHour *int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	if expiredHour != nil {
		expirationTime := time.Now().Add(time.Duration(*expiredHour) * time.Hour)
		token.Claims.(jwt.MapClaims)["exp"] = expirationTime.Unix()
	}

	return token.SignedString([]byte(service.SecretKey))
}

func (service *JWTService) GenerateRefreshToken(claims map[string]interface{}) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	expirationTime := time.Now().Add(time.Hour * 72).Unix()
	token.Claims.(jwt.MapClaims)["exp"] = expirationTime

	return token.SignedString([]byte(service.RefreshSecretKey))
}
