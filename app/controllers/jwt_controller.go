package controllers

import (
	"auth-microservice/app/services"
	"auth-microservice/models"
	u "auth-microservice/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {
	var body map[string]any

	if err := c.BodyParser(&body); err != nil {
		return u.ErrRes(c, fiber.StatusBadRequest, "body invalid")
	}

	expiredHour := 24

	if expHrs, ok := body["expired_hour"].(float64); ok {
		expiredHour = int(expHrs)
		if expiredHour <= 0 {
			return u.ErrRes(c, fiber.StatusBadRequest, "expired_hour must be greater than 0")
		}
	}

	delete(body, "expired_hour")
	services.GenerateNewToken(body, expiredHour)

	token, err := services.GenerateNewToken(body, expiredHour)
	if err != nil {
		return u.ErrRes(c, fiber.StatusInternalServerError, "could not generate token")
	}

	refresh, err := services.GenerateRefreshToken(body)
	if err != nil {
		return u.ErrRes(c, fiber.StatusInternalServerError, "could not generate refresh_token")
	}

	data := fiber.Map{"access_token": token, "refresh_token": refresh}

	return u.OkRes(c, data)
}

func VerifyToken(c *fiber.Ctx) error {
	var body *models.JwtVerify

	if err := c.BodyParser(&body); err != nil {
		return u.ErrRes(c, fiber.StatusBadRequest, "body invalid")
	}

	if body.Token == "" && body.RefreshToken == "" {
		return u.ErrRes(c, fiber.StatusBadRequest, "body invalid")
	}

	tokenString, secret := extractTokenFromBody(body)

	valid, claims, err := services.ValidateToken(tokenString, secret)

	if !valid {
		return u.ErrRes(c, fiber.StatusUnauthorized, err.Error())
	}

	return u.OkRes(c, claims)
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
