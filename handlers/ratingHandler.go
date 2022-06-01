package handlers

import (
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserRating(c *gin.Context) {
	var rating models.RatingRequest
	userI, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"message": "internal server get error"})
	}
	_, ok := userI.(*models.Buyer)
	if !ok {
		c.JSON(500, gin.H{"message": "invalid assert"})
	}
	err := c.BindJSON(&rating)
	if err != nil {
		c.JSON(404, gin.H{"message": "please fill al fields"})
		c.Abort()
		return
	}
	seller, err := h.DB.FindSellerById(rating.Id)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		c.Abort()
		return
	}

	seller.Rating = (seller.TotalRatings + rating.Rating) / seller.NumberOfRatingsReceived
}
