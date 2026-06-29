package handler

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Hello is the handler that answers to route /api/hello
func Hello(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello, world! You API is working! 🚀",
		Status:  200,
	}

	// We tell that the content is JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
