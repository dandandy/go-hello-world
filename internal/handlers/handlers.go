package handlers

import (
	"net/http"

	"github.com/dandandy/go-hello-world/internal/configuration"
	"github.com/dandandy/go-hello-world/internal/handlers/health"
	"github.com/dandandy/go-hello-world/internal/handlers/helloworld"
	"github.com/dandandy/go-hello-world/internal/handlers/metadata"
)

// Adds all of the individual handlers to the serve mux
func Add(config configuration.Bundle, serveMux *http.ServeMux) error {
	err := metadata.Add(config, serveMux)
	if err != nil {
		return err
	}

	helloworld.Add(serveMux)
	health.Add(serveMux)
	return nil
}
