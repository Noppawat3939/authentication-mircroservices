package controllers

import (
	"auth-microservice/app/services"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 400, "success": false, "message": "body invalid"})
	}

	expiredHour := 24

	if expHrs, ok := body["expired_hour"].(float64); ok {
		expiredHour = int(expHrs)
		if expiredHour <= 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 400, "success": false, "message": "expired_hour must be greater than 0"})
		}
	}

	delete(body, "expired_hour")
	services.GenerateNewToken(body, expiredHour)

	token, err := services.GenerateNewToken(body, expiredHour)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    500,
			"success": false,
			"message": "could not generate token",
		})
	}

	refresh, err := services.GenerateRefreshToken(body)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    500,
			"success": false,
			"message": "could not generate refresh_token",
		})
	}

	data := fiber.Map{"access_token": token, "refresh_token": refresh}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "success": true, "data": data})
}

func VerifyToken(c *fiber.Ctx) error {
	authorizeation := c.Get("Authorization")

	secretKey := os.Getenv("JWT_SECRET")

	valid, claims, err := services.ValidateToken(authorizeation, secretKey)

	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 401, "success": false, "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "success": true, "data": claims})
}
