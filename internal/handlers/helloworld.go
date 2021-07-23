package handlers

import "net/http"

const HelloWorldPath = "/"

var helloWorld = []byte(`hello world`)

func HelloWorld(rw http.ResponseWriter, req *http.Request) {
	rw.Write(helloWorld)
	rw.WriteHeader(200)
}
