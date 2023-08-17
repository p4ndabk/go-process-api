package main

import (
	"fmt"
	"github.com/p4ndabk/go-process-api/routes"
)

func main() {
	port := ":8080"
    fmt.Println("initial server on port", port)
	routes.HanddleRequests(port)
}