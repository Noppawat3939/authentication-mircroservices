package routes

import (
	"auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	r := router.Group("user")

	userController := controllers.NewUserController()

	r.Post("/register", userController.RegisterUser)
	r.Post("/login", userController.LoginUser)
}
