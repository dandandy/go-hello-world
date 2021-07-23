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

	err := addHandlers()
	if err != nil {
		log.Fatalf("failed to start with error %s", err)
	}

	log.Printf("listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func addHandlers() error {
	config, err := configuration.Load()
	if err != nil {
		return err
	}

	metadataHandler, err := handlers.Metadata(config)
	if err != nil {
		return err
	}

	http.HandleFunc(handlers.MetadataPath, metadataHandler)
	http.HandleFunc(handlers.HelloWorldPath, handlers.HelloWorld)
	http.HandleFunc(handlers.HealthCheckPath, handlers.HealthCheck)
	return nil
}
