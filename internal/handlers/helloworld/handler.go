package helloworld

import (
	"net/http"

	"github.com/dandandy/go-hello-world/internal/utils"
)

const path = "/"

var helloWorld = []byte(`hello world`)

func Add(s *http.ServeMux) {
	s.HandleFunc(path, handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	_, err := rw.Write(helloWorld)
	if err != nil {
		utils.InternalServerErrorResponse(rw)
		return
	}
}
