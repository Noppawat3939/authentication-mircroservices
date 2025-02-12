package controllers

import (
	r "auth-microservice/internal/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func EndpointNotFound(c *fiber.Ctx) error {
	path := c.Path()
	msg := fmt.Sprintf("endpoint %s not found", path)

	return r.Error(c, fiber.StatusNotFound, msg)
}
