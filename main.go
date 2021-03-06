package main

import (
	"log"
	"net/http"

	"./album"

	"github.com/gorilla/handlers"
)

func main() {
	router := album.NewRouter()

	//these two lines are important in order to allow access from the front-end side
	//to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	//launch server with CORS validations
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
