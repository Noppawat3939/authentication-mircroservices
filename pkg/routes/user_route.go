package routes

import (
	ctl "auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	r := router.Group("user")

	r.Post("/register", ctl.RegisterUser)
	r.Post("/login", ctl.LoginUser)
}
