package handlers

import (
	"net/http"

	//"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) AllBuyerOrders(c *gin.Context) {
	buyerId := c.Param("id")
	buyerIdStr, _ := strconv.Atoi(buyerId)

	buyerWithOrder := h.DB.GetAllBuyerOrder(uint(buyerIdStr))

	c.JSON(http.StatusFound, gin.H{
		"All Products": buyerWithOrder,
	})
}

func (h *Handler) AllSellerOrders(c *gin.Context)   {}
func (h *Handler) SellerTotalOrders(c *gin.Context) {}
