package routes

import (
	"auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func CheckRoute(r fiber.Router) {
	r.Get("/check", controllers.CheckServerRunning)
}
