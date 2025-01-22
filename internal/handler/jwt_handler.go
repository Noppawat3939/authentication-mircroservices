package handler

import (
	"auth-microservice/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type JWTHandler struct {
	Service *service.JWTService
}

func (h *JWTHandler) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Expired  *int64 `json:"expired_hour,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	token, err := h.Service.GenerateToken(payload.Email, payload.Username, payload.Expired)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating token: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"token": token, "success": true}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
