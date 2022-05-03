package handlers

import "github.com/gin-gonic/gin"

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
