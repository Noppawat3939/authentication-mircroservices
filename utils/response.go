package utils

import (
	"github.com/gofiber/fiber/v2"
)

type R map[string]interface{}

func OkRes(c *fiber.Ctx, data any) error {
	response := fiber.Map{"code": 200, "success": true}

	if data != nil {
		response["data"] = data
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func ErrRes(c *fiber.Ctx, statusCode int, message string) error {
	response := fiber.Map{"code": statusCode, "success": false, "message": message}

	return c.Status(fiber.StatusOK).JSON(response)
}
