package utils

import (
	"encoding/json"
	"net/http"
)

type R map[string]interface{}

func SuccessResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := R{
		"success":        true,
		"data":           data,
		"success_status": http.StatusOK,
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := R{
		"success":      false,
		"message":      message,
		"error_status": statusCode,
	}

	json.NewEncoder(w).Encode(response)
}
