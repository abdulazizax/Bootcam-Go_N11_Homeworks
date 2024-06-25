package main

import (
	"api-gateway/api"
	"log"
	"net/http"
	"os"
)

func main() {
	server := api.New()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, server); err != nil {
		log.Fatal("Failed to run HTTP server: ", err)
	}

	log.Print("Server stopped")
}
