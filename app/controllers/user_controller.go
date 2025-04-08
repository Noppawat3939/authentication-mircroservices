package controllers

import (
	s "auth-microservice/app/services"
	"auth-microservice/internal/response"
	r "auth-microservice/internal/response"
	"auth-microservice/models"
	"auth-microservice/repository"

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

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	data := fiber.Map{"user": newUser}

	return r.Success(c, data)
}

func (uc *UserController) LoginUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body invalid")
	}

	foundUser, err := uc.userService.LoginUser(&user)

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	payload := map[string]any{
		"id":    foundUser.ID,
		"email": foundUser.Email,
	}

	expiredHour := 24

	token, err := s.GenerateNewToken(payload, expiredHour)

	if err != nil {
		return r.Error(c, fiber.StatusInternalServerError, "could not generate token")

	}

	refresh, err := s.GenerateRefreshToken(payload)

	if err != nil {
		return r.Error(c, fiber.StatusInternalServerError, "could not generate refresh_token")
	}

	data := fiber.Map{"access_token": token, "refresh_token": refresh}

	return r.Success(c, data)
}
