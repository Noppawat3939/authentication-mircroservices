package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// jwtSecret := viper.GetString("app.jwt_secret")
	// jwtRefreshSecret := viper.GetString("app.jwt_refresh_secret")

	// jwtHandler := handler.NewJwtHandler(jwtSecret, jwtRefreshSecret)

	// r.HandleFunc("/jwt/generate", jwtHandler.GenerateTokenHandler)
	// r.HandleFunc("/jwt/verify", jwtHandler.Verify)

	// http.ListenAndServe(address, r)
	app.Listen(":8000")
}
