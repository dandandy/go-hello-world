package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dandandy/go-hello-world/internal/configuration"
)

func TestMetadata(t *testing.T) {
	t.Run("does not throw an error", func(t *testing.T) {
		wantErr := false
		got, err := Metadata(configuration.Bundle{})
		if (err != nil) != wantErr {
			t.Errorf("Metadata() error = %v, wantErr %v", err, wantErr)
			return
		}

		// check function returned is not nil
		if got == nil {
			t.Error("Metadata() = nil, want not nil")
		}
	})
}

func Test_metadata_Handler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, MetadataPath, nil)
	if err != nil {
		t.Error(err)
	}

	m := metadata{
		response: []byte(`foo`),
	}
	handler := http.HandlerFunc(m.Handler)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `foo`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	if rr.Header().Get(contentType) != applicationJson {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			rr.Header().Get(contentType), applicationJson)
	}
}
