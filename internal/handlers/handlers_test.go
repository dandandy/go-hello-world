package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dandandy/go-hello-world/internal/configuration"
)

func TestAdd(t *testing.T) {
	serveMux := http.NewServeMux()

	t.Run("it successfully adds handlers to serve mux", func(t *testing.T) {
		wantErr := false
		config := configuration.Bundle{}
		if err := Add(config, serveMux); (err != nil) != wantErr {
			t.Errorf("Add() error = %v, wantErr %v", err, wantErr)
		}
	})

	t.Run("serve mux has expected number of handlers added", func(t *testing.T) {
		serveMuxValue := reflect.ValueOf(&serveMux).Elem()

		expectedNumberOfHandlers := 3
		if handlers := reflect.Indirect(serveMuxValue).FieldByName("m"); handlers.Len() != expectedNumberOfHandlers {
			t.Errorf("Add() added %v handlers to serve mux, expected %v", handlers.Len(), expectedNumberOfHandlers)
		}
	})

	t.Run("test server mux starts and responds with success for basic request", func(t *testing.T) {
		respRec := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Error("Creating request failed")
		}

		serveMux.ServeHTTP(respRec, req)

		expect := 200
		if got := respRec.Code; got != expect {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})
}
