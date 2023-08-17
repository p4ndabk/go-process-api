package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("Home Page!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("is alive!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}