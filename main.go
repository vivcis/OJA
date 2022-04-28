package main

import (
	"fmt"
	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	fmt.Println("Oja")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.SetupDB()
	router, port := router.SetupRouter()

	router.Run(port)

}
