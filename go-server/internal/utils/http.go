package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// SuccessResponse is a helper that responds with a 2** status code
func SuccessResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

// ErrorResponse is a helper that respond with 4** or 5** status code
func ErrorResponse(w http.ResponseWriter, msg string, statusCode int) {
	response := map[string]string{"error": msg}

	w.Header().Add("Content-Type", "application/json")

	resp, err := json.Marshal(response)
	if err != nil {
		slog.Error("failed to marshal response", "err", err)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(resp)
}
