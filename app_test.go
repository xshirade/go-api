package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "Gorilla!"
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", res.Body.String(), expected)
	}
}
