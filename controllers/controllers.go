package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "is aline!")
}