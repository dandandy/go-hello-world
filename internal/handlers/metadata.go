package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dandandy/go-hello-world/internal/configuration"
)

const MetadataPath = "/metadata"

type metadata struct {
	response []byte
}

type metadataResponseBody struct {
	ApplicationName string `json:"name"`
	Version         string `json:"version"`
	LastCommitSha   string `json:"lastCommitSha"`
	Description     string `json:"description"`
}

func Metadata(c configuration.Bundle) (func(http.ResponseWriter, *http.Request), error) {
	response := metadataResponseBody{
		ApplicationName: c.GetApplicationName(),
		Version:         c.GetVersion(),
		LastCommitSha:   c.GetLastCommitSha(),
		Description:     c.GetDescription(),
	}

	responseJson, err := response.toJson()
	if err != nil {
		return nil, err
	}

	return metadata{
		response: responseJson,
	}.Handler, nil
}

func (m *metadataResponseBody) toJson() ([]byte, error) {
	return json.Marshal(m)
}

func (m metadata) Handler(rw http.ResponseWriter, req *http.Request) {
	contentTypeApplicationJson(rw.Header())
	rw.WriteHeader(200)
	rw.Write(m.response)
}
