package routes

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/p4ndabk/go-process-api/controllers"
)


func HanddleRequests(port string) {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/", controllers.Home).Methods("GET")
	rotas.HandleFunc("/health", controllers.Health).Methods("GET")
	rotas.HandleFunc("/process", controllers.Process).Methods("POST")
	rotas.HandleFunc("/go-process", controllers.GoProcess).Methods("POST")
	
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))
}