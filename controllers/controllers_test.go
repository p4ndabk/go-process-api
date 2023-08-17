package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code %d but got %d", http.StatusOK, rr.Code)

	expected := `"Home Page!"`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body mismatch")
}
 func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code %d but got %d", http.StatusOK, rr.Code)

	expected := `"is alive!"`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body mismatch")
}




