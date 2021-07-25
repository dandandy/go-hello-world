package helloworld

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
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

	expected := `hello world`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAdd(t *testing.T) {
	t.Run("it handles requests to path when handler added to serve mux", func(t *testing.T) {
		serveMux := http.NewServeMux()
		respRec := httptest.NewRecorder()
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Error("Creating request failed")
		}

		Add(serveMux)
		serveMux.ServeHTTP(respRec, req)

		expectedStatusCode := 200
		if got := respRec.Code; got != expectedStatusCode {
			t.Errorf("expected %v, got %v", expectedStatusCode, got)
		}

		expectedBody := string(helloWorld)
		if got := respRec.Body.String(); got != expectedBody {
			t.Errorf("expected %v, got %v", expectedBody, got)
		}
	})
}
