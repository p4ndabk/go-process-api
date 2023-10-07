package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
)

type City struct {
	CodeIBGE int `json:"code_ibge"`
	CodeUF int `json:"code_uf"`
	Name string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Capital int `json:"capital"`
}

type Cities struct {
	Cities []City `json:"cities"`
}


func CityList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")	
	
	jsonData, err := os.Open("storage/cities.json")
	if err != nil {
		log.Println("error open file cities.json")
	}

	var cityData Cities	
	decoder := json.NewDecoder(jsonData)
	if err := decoder.Decode(&cityData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cityData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer jsonData.Close()
}
