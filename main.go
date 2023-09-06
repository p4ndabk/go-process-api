package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/p4ndabk/go-process-api/mylog"
	"github.com/p4ndabk/go-process-api/routes"
)

func main() {
	loadEnv()
	logfile := mylog.InitLog()
    log.Println("Starting the application...")

	routes.HanddleRequests(os.Getenv("SERVER_PORT"))
	defer logfile.Close()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
