package main

import (
	"auth-microservice/internal/handler"
	"auth-microservice/internal/router"
	"auth-microservice/internal/service"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func initializeConfig() {
	viper.SetDefault("app.port", 8000)
	viper.AutomaticEnv()
}

func main() {
	initializeConfig()

	port := viper.GetInt("app.port")
	address := fmt.Sprintf(":%d", port)

	jwtSecret := viper.GetString("app.jwt_secret")

	fmt.Print(1, jwtSecret)

	jwtService := &service.JWTService{
		SecretKey: "your-secret-key",
	}

	jwtHandler := &handler.JWTHandler{
		Service: jwtService,
	}

	r := router.JwtRouter(jwtHandler)

	fmt.Printf("🚀 Server is running on port %d\n", port)
	http.ListenAndServe(address, r)
}
