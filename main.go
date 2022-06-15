package main

import (
	"github.com/joho/godotenv"
	"log"
	"tiny-url/bootstrap"
)

func main() {
	_ = godotenv.Load()
	err := bootstrap.RootApp.Execute()
	if err != nil {
		log.Fatal("Server Can't started")
	}
}
