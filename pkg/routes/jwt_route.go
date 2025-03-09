package routes

import (
	ctl "auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func JwtRoutes(router fiber.Router) {
	r := router.Group("jwt")

	r.Post("/generate", ctl.GetJwtToken)
	r.Post("/verify", ctl.VerifyToken)
}
