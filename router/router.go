package router

import (
	"log"
	"os"

	"github.com/decadevs/shoparena/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() (*gin.Engine, string) {
	router := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apirouter := router.Group("/api/v1")

	apirouter.GET("/ping", handlers.PingHandler)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
