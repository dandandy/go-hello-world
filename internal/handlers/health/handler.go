package health

import (
	"encoding/json"
	"net/http"

	"github.com/dandandy/go-hello-world/internal/utils"
)

const path = "/health"

type response struct {
	Healthy      bool         `json:"healthy"`
	Dependencies []dependency `json:"dependencies"`
}

type dependency struct {
	Name    string `json:"name"`
	Healthy bool   `json:"healthy"`
}

func Add(s *http.ServeMux) {
	s.HandleFunc(path, handler)
}

// Lightweight endpoint to show that the application is alive and responsive.
func handler(rw http.ResponseWriter, req *http.Request) {
	resp := response{
		Healthy:      true,
		Dependencies: []dependency{},
	}

	responseJson, err := resp.toJson()
	if err != nil {
		utils.InternalServerErrorResponse(rw)
		return
	}
	utils.ContentTypeApplicationJson(rw.Header())
	rw.WriteHeader(200)
	_, err = rw.Write(responseJson)
	if err != nil {
		utils.InternalServerErrorResponse(rw)
		return
	}
}

func (h *response) toJson() ([]byte, error) {
	return json.Marshal(h)
}
