package routes

import (
	"auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func JwtRoutes(router fiber.Router) {
	r := router.Group("jwt")

	r.Post("/generate", controllers.GetJwtToken)
	r.Post("/verify", controllers.VerifyToken)
}
