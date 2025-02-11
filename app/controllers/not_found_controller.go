package controllers

import (
	u "auth-microservice/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func EndpointNotFound(c *fiber.Ctx) error {
	path := c.Path()

	message := fmt.Sprintf("endpoint %s not found", path)
	return u.ErrRes(c, fiber.StatusNotFound, message)
}
