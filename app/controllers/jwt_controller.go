package controllers

import (
	"auth-microservice/app/services"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 400, "success": false, "message": "body invalid"})
	}

	expiredHour := 24

	hash := sha256.New()

	refresh := "refresh_token_secret_key" + time.Now().String()

	_, err := hash.Write([]byte(refresh))

	if err != nil {
		return err
	}

	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(3)).Unix())
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	if expHrs, ok := body["expired_hour"].(float64); ok {
		expiredHour = int(expHrs)
		if expiredHour <= 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 400, "success": false, "message": "expired_hour must be greater than 0"})
		}
	}

	services.GenerateNewToken(body, expiredHour)

	token, err := services.GenerateNewToken(body, expiredHour)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    500,
			"success": false,
			"message": "could not generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "success": true, "data": fiber.Map{"access_token": token, "t": t}})
}
