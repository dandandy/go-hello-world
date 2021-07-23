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

	serveMux := http.NewServeMux()
	err := addHandlers(serveMux)
	if err != nil {
		log.Fatalf("failed to start with error %s", err)
	}

	log.Printf("listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, serveMux))
}

func addHandlers(serveMux *http.ServeMux) error {
	config, err := configuration.Load()
	if err != nil {
		return err
	}

	err = handlers.AddMetadataHandler(config, serveMux)
	if err != nil {
		return err
	}

	handlers.AddHelloWorldHandler(serveMux)
	handlers.AddHealthCheckHandler(serveMux)
	return nil
}
