package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jalawz/library-api/internal/handler"
)

func main() {
	// Register our route
	http.HandleFunc("/api/hello", handler.Hello)

	// Friendly terminal message
	fmt.Println("🚀 Server running on http://localhost:8080")
	fmt.Println("📡 Endpoint: http://localhost:8080/api/hello")

	// Iniate server in port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
