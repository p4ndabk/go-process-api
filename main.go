package main

import (
	"github.com/p4ndabk/go-process-api/routes"
)

func main() {
	port := ":8080"
	routes.HanddleRequests(port)
}