package handlers

import (
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) CreateProducts(c *gin.Context) {
	sellerId := c.Param("sellerid")
	sellerIdInt, _ := strconv.Atoi(sellerId)

	product := models.Product{}

	log.Println("here 1", product.SellerId)
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to bind product",
		})
		return
	}

	product.SellerId = uint(sellerIdInt)

	err = h.DB.CreateProduct(product)
	if err != nil {
		log.Println("check error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating product in handler"})
		return
	}

	response := gin.H{

		"data":      product,
		"status":    http.StatusText(http.StatusCreated),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, response)
}
