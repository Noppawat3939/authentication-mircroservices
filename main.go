package main

import (
	"auth-microservice/pkg/middleware"
	"auth-microservice/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	api := app.Group("/api/v1")

	routes.JwtRoutes(app)
	routes.CheckRoute(api)
	routes.NotFoundRoute(app)

	app.Listen(":8000")
}
