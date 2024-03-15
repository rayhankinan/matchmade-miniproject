package main

import (
	"log"
	"net/http"

	"service/pkg/config"
	"service/pkg/db"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db.Connect()

	// Add the routes

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
