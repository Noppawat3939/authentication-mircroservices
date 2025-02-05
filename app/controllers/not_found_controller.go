package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func EndpointNotFound(c *fiber.Ctx) error {
	path := c.Path()
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": fmt.Sprintf("endpoint %s not found", path),
	})
}
