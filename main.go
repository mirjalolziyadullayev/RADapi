package main

import (
	"fmt"
	"net/http"
	"RADserver/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
// Create a new router
r := mux.NewRouter()

// Define your API routes
r.HandleFunc("/users", handler.UsersHandler).Methods("POST","GET","DELETE")

// Define the allowed origins, methods, and headers
allowedOrigins := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500"})
allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})

// Enable CORS with the specified options
corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)

// Start the server
fmt.Printf("Server is running on port :8080")
http.Handle("/", corsHandler)
http.ListenAndServe(":8080", nil)
}