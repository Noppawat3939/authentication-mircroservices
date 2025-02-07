package main

import (
	"auth-microservice/pkg/middleware"
	"auth-microservice/pkg/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	middleware.FiberMiddleware(app)

	api := app.Group("/api/v1")

	routes.JwtRoutes(api)
	routes.CheckRoute(api)
	routes.NotFoundRoute(app)

	app.Listen(":8000")
}
