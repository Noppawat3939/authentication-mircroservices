package routes

import "github.com/gofiber/fiber/v2"

func JwtRoutes(a *fiber.App) {
	route := a.Group("/api/v1/jwt")

	route.Post("/generate", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
	})
}
