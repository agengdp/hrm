package main

import (
	"anwarmedika/hrm/configs"
	"anwarmedika/hrm/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server.Init()
	configs.DBInit()

}
