package handlers

import (
	"log"
	"net/http"

	//"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) AllBuyerOrders(c *gin.Context) {
	buyerId := c.Param("id")
	buyerIdStr, _ := strconv.Atoi(buyerId)

	buyerWithOrder, err := h.DB.GetAllBuyerOrder(uint(buyerIdStr))
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusFound, gin.H{
		"Orders": buyerWithOrder,
	})
}

func (h *Handler) AllSellerOrders(c *gin.Context)   {}
func (h *Handler) SellerTotalOrders(c *gin.Context) {}
