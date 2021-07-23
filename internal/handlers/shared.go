package handlers

import "net/http"

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
)

func contentTypeApplicationJson(header http.Header) {
	header.Add(contentType, applicationJson)
}
