package utils

import (
	"net/http"

	"log"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	internalServerErrorBody = []byte(`something went wrong`)
)

func ContentTypeApplicationJson(header http.Header) {
	header.Add(ContentType, ApplicationJson)
}

func InternalServerErrorResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	_, err := rw.Write(internalServerErrorBody)
	if err != nil {
		log.Default().Printf("failed to write internal server error response with error %s", err)
	}
}
