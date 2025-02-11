package controllers

import (
	u "auth-microservice/utils"

	"github.com/gofiber/fiber/v2"
)

func CheckServerRunning(c *fiber.Ctx) error {
	data := fiber.Map{"message": "server is running"}

	return u.OkRes(c, data)
}
