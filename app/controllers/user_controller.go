package controllers

import (
	r "auth-microservice/internal/response"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	data := fiber.Map{"message": "register user success"}

	return r.Success(c, data)
}

func LoginUser(c *fiber.Ctx) error {
	return r.Success(c, nil)
}
