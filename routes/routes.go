package routes

import (
	"net/http" //go's standard library for HTTP handling
	"github.com/KumarAyushh/go-backend-playground/internal/handlers"
)

// SetupRoutes initializes the routes for the application and returns a ServeMux
func SetupRoutes() *http.ServeMux {

	// Create a new ServeMux which will act as our router
	mux := http.NewServeMux()

	//It is a handler function that will be called when a request is made to the "/hello" endpoint.
	mux.HandleFunc("/hello", handlers.HelloHandler)

	// Additional routes can be added here in the future

	// Ab poora configured router return kar rahe hain 
	// Ye router main.go me use hoga: 
	// http.ListenAndServe(":8080", router)
	return mux
}