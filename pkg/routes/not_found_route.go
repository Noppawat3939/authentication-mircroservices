package routes

import (
	"auth-microservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(a *fiber.App) {
	a.Use(controllers.EndpointNotFound)
}
