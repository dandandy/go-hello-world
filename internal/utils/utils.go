package utils

import "net/http"

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func ContentTypeApplicationJson(header http.Header) {
	header.Add(ContentType, ApplicationJson)
}
