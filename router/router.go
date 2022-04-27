package router

import (
	"github.com/decadevs/multivendor/handlers"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupRouter() (*gin.Engine, string) {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/ping", handlers.PingHandler)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
