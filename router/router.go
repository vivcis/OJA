package router

import (
	"log"
	"net/http"
	"os"

	"github.com/decadevs/shoparena/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Router struct {
	ContentType string
	handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

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
