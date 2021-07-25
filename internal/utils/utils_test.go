package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInternalServerErrorResponse(t *testing.T) {
	t.Run("it adds correct body and status code to response writer struct", func(t *testing.T) {
		rw := httptest.NewRecorder()
		InternalServerErrorResponse(rw)

		if gotCode := rw.Result().StatusCode; gotCode != http.StatusInternalServerError {
			t.Errorf("expected code = %v, got %v", http.StatusInternalServerError, gotCode)
		}

		if rw.Body.String() != string(internalServerErrorBody) {
			t.Errorf("expected rw.Body = %v, got %v", string(internalServerErrorBody), rw.Body.String())
		}
	})
}

func TestContentTypeApplicationJson(t *testing.T) {
	t.Run("adds content type app json to header", func(t *testing.T) {
		header := http.Header{}
		ContentTypeApplicationJson(header)

		if header.Get(ContentType) != ApplicationJson {
			t.Errorf("content type application json was not added to header")
		}
	})
}
