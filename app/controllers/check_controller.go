package controllers

import "github.com/gofiber/fiber/v2"

func CheckServerRunning(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "server is running"})
}
