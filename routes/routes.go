package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/go-process-api/controllers"
	"log"
	"net/http"
)

func HanddleRequests(port string) {
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/", controllers.Home).Methods("GET")
	routes.HandleFunc("/health", controllers.Health).Methods("GET")
	routes.HandleFunc("/process", controllers.Process).Methods("POST")
	routes.HandleFunc("/go-process", controllers.GoProcess).Methods("POST")

	//solar
	routes.HandleFunc("/solar", controllers.Solar).Methods("GET")
	routes.HandleFunc("/city", controllers.CityList).Methods("GET")


	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, routes))
}
