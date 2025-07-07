package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/abdullahnettoor/tastybites/internal/api/dto"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := dto.CommonResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	WriteJSONResponse(w, statusCode, response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := dto.CommonResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
	WriteJSONResponse(w, statusCode, response)
}

func GetUserIDFromContext(r *http.Request) (int, error) {
	userID, ok := r.Context().Value("userId").(int)
	if !ok {
		return 0, errors.New("userID not found in context")
	}
	return userID, nil
}
