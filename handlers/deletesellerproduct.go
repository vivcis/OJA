package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteSellerProduct(c *gin.Context) {
	productID := c.Param("id")
	productIdInt, _ := strconv.Atoi(productID)

	err := h.DB.DeleteProduct(uint(productIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting product"})
		return
	}

	c.JSON(http.StatusOK, "successfully deleted")
}
