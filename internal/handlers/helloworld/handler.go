package helloworld

import "net/http"

const path = "/"

var helloWorld = []byte(`hello world`)

func Add(s *http.ServeMux) {
	s.HandleFunc(path, handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	rw.Write(helloWorld)
	rw.WriteHeader(200)
}
