package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dandandy/go-hello-world/internal/configuration"
)

const metadataPath = "/metadata"

type metadataResponse struct {
	response []byte
}

type metadataResponseBody struct {
	ApplicationName string `json:"name"`
	Version         string `json:"version"`
	LastCommitSha   string `json:"lastCommitSha"`
	Description     string `json:"description"`
}

func AddMetadataHandler(c configuration.Bundle, s *http.ServeMux) error {
	handler, err := newMetadataHandler(c)
	if err != nil {
		return err
	}
	s.HandleFunc(metadataPath, handler)
	return nil
}

func newMetadataHandler(c configuration.Bundle) (func(http.ResponseWriter, *http.Request), error) {
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

	return metadataResponse{
		response: responseJson,
	}.handler, nil
}

func (m *metadataResponseBody) toJson() ([]byte, error) {
	return json.Marshal(m)
}

func (m metadataResponse) handler(rw http.ResponseWriter, req *http.Request) {
	contentTypeApplicationJson(rw.Header())
	rw.WriteHeader(200)
	rw.Write(m.response)
}
