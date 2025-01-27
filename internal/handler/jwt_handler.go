package handler

import (
	"auth-microservice/internal/service"
	"auth-microservice/utils"
	"encoding/json"
	"net/http"
)

type JWTHandler struct {
	Service *service.JWTService
}

func NewJwtHandler(secretKey string) *JWTHandler {
	return &JWTHandler{Service: service.NewJwtService(secretKey)}
}

func (h *JWTHandler) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, "invalid request method", http.StatusMethodNotAllowed)

		return
	}

	var payload map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.ErrorResponse(w, "body should be object value", http.StatusBadRequest)
		return
	}

	var expired *int64

	if val, ok := payload["expired_hour"]; ok {
		if hours, ok := val.(float64); ok {
			exp := int64(hours)
			expired = &exp
		}
		delete(payload, "expired_hour")
	}

	token, err := h.Service.GenerateToken(payload, expired)
	if err != nil {
		utils.ErrorResponse(w, "failed generate token", http.StatusInternalServerError)

	} else {
		data := map[string]interface{}{"token": token}
		utils.SuccessResponse(w, data)
	}
}
