package main

import (
	"log"
	"net/http"

	"github.com/dandandy/go-hello-world/internal/configuration"
	"github.com/dandandy/go-hello-world/internal/handlers"
)

var PORT = ":8080"

func main() {
	log.Print("Starting server...")

	config := configuration.Load()

	serveMux := http.NewServeMux()
	err := handlers.Add(config, serveMux)
	if err != nil {
		log.Fatalf("failed to start with error %s", err)
	}

	log.Printf("listening on port http://localhost%s", PORT)
	log.Fatal(http.ListenAndServe(PORT, serveMux))
}
