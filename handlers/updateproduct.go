package handlers

import (
	"fmt"
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//UPDATE PRODUCT BY ID
func (h *Handler) UpdateProduct(c *gin.Context) {
	prod := c.Param("id")
	prodID, _ := strconv.Atoi(prod)
	prodIdUint := uint(prodID)

	var product = models.Product{}

	if err := c.BindJSON(&product); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error binding in updateProduct handler"})
		return
	}

	product.ID = prodIdUint

	fmt.Println(product.Title, product.Price)

	err := h.DB.UpdateProductByID(product)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating product by ID"})
		return
	}

	c.JSON(http.StatusOK, "updated successfully")
}
