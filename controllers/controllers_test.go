package controllers

import (
	"bytes"
	"encoding/json"
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

func TestProcess(t *testing.T) {
	var body = RequestProcessBody{Process: 1, Worker: 1}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/process", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Process)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code %d but got %d", http.StatusOK, rr.Code)
	
	expected := `{"process_0":true}`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body mismatch")
}

func TestGoProcess(t *testing.T) {
	var body = RequestProcessBody{Process: 1, Worker: 1}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/go-process", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GoProcess)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code %d but got %d", http.StatusOK, rr.Code)
	expected := `{"process_0":true}`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body mismatch")
}



