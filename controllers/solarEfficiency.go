package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Solar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Solar Page Efficiency!"}
	fmt.Println(response)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


	