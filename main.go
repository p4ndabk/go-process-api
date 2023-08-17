package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/p4ndabk/go-process-api/routes"
)

func main() {
	loadEnv()
	routes.HanddleRequests(os.Getenv("SERVER_PORT"))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}