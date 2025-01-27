package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"success": true,
	}

	if data != nil {
		response["data"] = data
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"success":      false,
		"message":      message,
		"error_status": statusCode,
	}

	json.NewEncoder(w).Encode(response)
}
