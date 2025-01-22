package router

import (
	"auth-microservice/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func JwtRouter(jwtHandler *handler.JWTHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/jwt/generate", jwtHandler.GenerateTokenHandler).Methods(http.MethodPost)
	return r
}
