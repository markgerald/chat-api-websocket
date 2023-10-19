package main

import (
	"github.com/joho/godotenv"
	"github.com/markgerald/chat-api-challenge/database"
	"github.com/markgerald/chat-api-challenge/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DbConnect()
	routes.HandleRequests()
}
