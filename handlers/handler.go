package handlers

import (
	"github.com/decadevs/shoparena/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	DB database.DB
}

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) SearchDBQuery(c *gin.Context) {
	products, err := h.DB.SearchDB(c)
	if err != nil {
		log.Println("handler error in search function", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, products)
}
