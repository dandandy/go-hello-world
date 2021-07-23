package handlers

import "net/http"

const helloWorldPath = "/"

var helloWorld = []byte(`hello world`)

func AddHelloWorldHandler(s *http.ServeMux) {
	s.HandleFunc(helloWorldPath, helloWorldHandler)
}

func helloWorldHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(helloWorld)
	rw.WriteHeader(200)
}
