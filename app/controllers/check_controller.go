package controllers

import (
	r "auth-microservice/internal/response"

	"github.com/gofiber/fiber/v2"
)

func CheckServerRunning(c *fiber.Ctx) error {
	data := fiber.Map{"message": "server is running"}

	return r.Success(c, data)
}
