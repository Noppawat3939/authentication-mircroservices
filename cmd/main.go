package main

import (
	"auth-microservice/internal/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func initializeConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func main() {
	initializeConfig()

	port := viper.GetInt("app.port")

	jwtSecret := viper.GetString("app.jwt_secret")
	jwtRefreshSecret := viper.GetString("app.jwt_refresh_secret")

	address := fmt.Sprintf(":%d", port)

	jwtHandler := handler.NewJwtHandler(jwtSecret, jwtRefreshSecret)

	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/jwt/generate", jwtHandler.GenerateTokenHandler)

	fmt.Printf("🚀 Server is running port %d\n", port)
	http.ListenAndServe(address, r)
}
