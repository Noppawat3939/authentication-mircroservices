package controllers

import (
	s "auth-microservice/app/services"
	r "auth-microservice/internal/response"
	"auth-microservice/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {
	var body map[string]any

	if err := c.BodyParser(&body); err != nil {
		return r.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	expiredHour := 24

	if expHrs, ok := body["expired_hour"].(float64); ok {
		expiredHour = int(expHrs)
		if expiredHour <= 0 {
			return r.Error(c, fiber.StatusBadRequest, "expired_hour must be greater than 0")
		}
	}

	delete(body, "expired_hour")
	s.GenerateNewToken(body, expiredHour)

	token, err := s.GenerateNewToken(body, expiredHour)
	if err != nil {
		return r.Error(c, fiber.StatusInternalServerError, "could not generate token")
	}

	refresh, err := s.GenerateRefreshToken(body)
	if err != nil {
		return r.Error(c, fiber.StatusInternalServerError, "could not generate refresh_token")
	}

	data := fiber.Map{"access_token": token, "refresh_token": refresh}

	return r.Success(c, data)
}

func VerifyToken(c *fiber.Ctx) error {
	var body *models.JwtVerify

	if err := c.BodyParser(&body); err != nil {
		return r.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	if body.Token == "" && body.RefreshToken == "" {
		return r.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	tokenString, secret := extractTokenFromBody(body)

	valid, claims, err := s.ValidateToken(tokenString, secret)

	if !valid {
		return r.Error(c, fiber.StatusUnauthorized, err)
	}

	return r.Success(c, claims)
}

func extractTokenFromBody(b *models.JwtVerify) (string, string) {
	var tokenString string
	var secretKey string

	if b.Token != "" {
		tokenString = b.Token
		secretKey = os.Getenv("JWT_SECRET")
	}

	if b.RefreshToken != "" {
		tokenString = b.RefreshToken
		secretKey = os.Getenv("JWT_REFRESH_SECRET")
	}

	return tokenString, secretKey
}
