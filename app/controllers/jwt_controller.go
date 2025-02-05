package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetJwtToken(c *fiber.Ctx) error {
	var body map[string]interface{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 400, "success": false, "message": "body invalid"})

	}
	fmt.Print(body)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"code": 200, "success": true})
}
