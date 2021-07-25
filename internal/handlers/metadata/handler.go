package metadata

import (
	"encoding/json"
	"net/http"

	"github.com/dandandy/go-hello-world/internal/configuration"
	"github.com/dandandy/go-hello-world/internal/utils"
)

const path = "/metadata"

type response struct {
	response []byte
}

type responseJsonBody struct {
	ApplicationName string `json:"name"`
	Version         string `json:"version"`
	LastCommitSha   string `json:"lastCommitSha"`
	Description     string `json:"description"`
}

func Add(c configuration.Bundle, s *http.ServeMux) error {
	handler, err := newHandler(c)
	if err != nil {
		return err
	}
	s.HandleFunc(path, handler)
	return nil
}

func newHandler(c configuration.Bundle) (func(http.ResponseWriter, *http.Request), error) {
	resp := responseJsonBody{
		ApplicationName: c.GetApplicationName(),
		Version:         c.GetVersion(),
		LastCommitSha:   c.GetLastCommitSha(),
		Description:     c.GetDescription(),
	}

	responseJson, err := resp.toJson()
	if err != nil {
		return nil, err
	}

	return response{
		response: responseJson,
	}.handler, nil
}

func (m *responseJsonBody) toJson() ([]byte, error) {
	return json.Marshal(m)
}

func (m response) handler(rw http.ResponseWriter, req *http.Request) {
	utils.ContentTypeApplicationJson(rw.Header())
	rw.WriteHeader(200)
	_, err := rw.Write(m.response)
	if err != nil {
		utils.InternalServerErrorResponse(rw)
		return
	}
}
