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
	user, ok := userI.(*models.Seller)
	if !ok {
		c.JSON(500, gin.H{"message": "invalid assert"})
	}
	seller, err := h.DB.FindSellerByEmail(user.Email)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		c.Abort()
		return
	}
	err = c.BindJSON(&rating)
	if err != nil {
		c.JSON(404, gin.H{"message": "something went wrong!"})
		c.Abort()
		return
	}
	seller.
	seller.Rating = (seller.TotalRatings + rating.Rating)/seller.NumberOfRatingsReceived)
}
