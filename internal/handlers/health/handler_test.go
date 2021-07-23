package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dandandy/go-hello-world/internal/utils"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)

	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"healthy":true,"dependencies":[]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if rr.Header().Get(utils.ContentType) != utils.ApplicationJson {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			rr.Header().Get(utils.ContentType), utils.ApplicationJson)
	}
}
