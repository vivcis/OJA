package handlers

import (
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProducts(c *gin.Context) {
	sellerId := c.Param("sellerid")
	sellerIdInt, _ := strconv.Atoi(sellerId)

	product := models.Product{}

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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error creating product in handler"})
		return
	}

	c.JSON(http.StatusCreated, "product created successfully")
}
