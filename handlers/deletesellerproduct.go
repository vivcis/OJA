package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func (h *Handler) GetUserFromContext(c *gin.Context) {
//	useI, ok := c.Get("user")
//}
func (h *Handler) DeleteSellerProduct(c *gin.Context) {
	seller, err := h.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
		return
	}

	productID := c.Param("id")
	productIdInt, _ := strconv.Atoi(productID)

	err = h.DB.DeleteProduct(uint(productIdInt), seller.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting product"})
		return
	}

	c.JSON(http.StatusOK, "successfully deleted")
}
