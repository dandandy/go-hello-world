package handlers

import (
	"encoding/json"
	"net/http"
)

const healthCheckPath = "/health"

type healthCheckResponse struct {
	Healthy      bool         `json:"healthy"`
	Dependencies []dependency `json:"dependencies"`
}

type dependency struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

func AddHealthCheckHandler(s *http.ServeMux) {
	s.HandleFunc(healthCheckPath, healthCheck)
}

// Lightweight endpoint to show that the application is alive and responsive.
func healthCheck(rw http.ResponseWriter, req *http.Request) {
	response := healthCheckResponse{
		Healthy:      true,
		Dependencies: []dependency{},
	}

	responseJson, err := response.toJson()
	if err != nil {
		rw.Write([]byte(`something went wrong`))
		rw.WriteHeader(500)
		return
	}
	contentTypeApplicationJson(rw.Header())
	rw.Write(responseJson)
	rw.WriteHeader(200)
}

func (h *healthCheckResponse) toJson() ([]byte, error) {
	return json.Marshal(h)
}
