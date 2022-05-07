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

func (h *Handler) SearchProductHandler(c *gin.Context) {
	//Equivalent to param
	category := c.Query("category")
	lowerPrice := c.Query("lower-price")
	upperPrice := c.Query("upper-price")
	name := c.Query("name")

	product, err := h.DB.SearchProduct(lowerPrice, upperPrice, category, name)
	if err != nil {
		log.Println("handler error in search product", err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(product) == 0 {
		c.JSON(http.StatusInternalServerError, "no such product")
		return
	}

	c.JSON(http.StatusFound, product)
}
