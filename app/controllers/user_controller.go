package controllers

import (
	s "auth-microservice/app/services"
	"auth-microservice/internal/response"
	r "auth-microservice/internal/response"
	"auth-microservice/models"
	"auth-microservice/repository"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *s.UserService
}

func NewUserController() *UserController {
	userRepo := repository.NewUserRepository()
	return &UserController{userService: s.NewUserService(userRepo)}
}

func (uc *UserController) RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	if user.Password == "" {
		return response.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	newUser, err := uc.userService.RegisterUser(&user)

	fmt.Print(111, newUser)

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	data := fiber.Map{"user": newUser}

	return r.Success(c, data)
}

func (uc *UserController) LoginUser(c *fiber.Ctx) error {
	return r.Success(c, nil)
}
