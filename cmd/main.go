package main

import (
	"net/http"
	"fmt"
	"github.com/KumarAyushh/go-backend-playground/routes"

)



func main(){
	//create a new ServeMux and register the routes
	mux := routes.SetupRoutes()

	fmt.Println("Server is running on port 8080...")

	//start the server on port 8080 and pass the custom handler
	http.ListenAndServe(":8080", mux)

	

}