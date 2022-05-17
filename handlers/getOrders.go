package handlers

import (
	"github.com/decadevs/shoparena/models"
	"net/http"

	//"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AllBuyerOrders(c *gin.Context) {
	user, present := c.Get("user")
	if !present {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "you are not logged in"})
	}

	buyer := user.(*models.Buyer)

	buyerWithOrder, err := h.DB.GetAllBuyerOrder(buyer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error getting order(s)",
		})
	}

	c.JSON(http.StatusFound, gin.H{
		"Buyer Order": buyerWithOrder,
	})
}

func (h *Handler) AllSellerOrders(c *gin.Context)   {}
func (h *Handler) SellerTotalOrders(c *gin.Context) {}
