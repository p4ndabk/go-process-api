package routes

import (
	"net/http"
	"github.com/p4ndabk/go-process-api/controllers"
)


func HanddleRequests(port string) {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/health", controllers.Health)
	http.ListenAndServe(port, nil)
}