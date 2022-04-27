package main

import (
	"fmt"
	"github.com/decadevs/multivendor/database"
	"github.com/decadevs/multivendor/handlers"
	"github.com/gin-gonic/gin"
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

	//default router
	router := gin.Default()

	router.GET("/ping", handlers.HandlerWelcomeStatement)

	//router.LoadHTMLGlob("templates/*.html")

	port := ":8085"
	router.Run(port)
}
