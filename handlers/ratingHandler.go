package handlers

import (
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SellerRating(c *gin.Context) {
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
		c.JSON(500, gin.H{"message": "Something Went Wrong!"})
		c.Abort()
		return
	}
	seller.TotalRatings += rating.Rating
	seller.NumberOfRatingsReceived += 1
	seller.Rating = (seller.TotalRatings + rating.Rating) / seller.NumberOfRatingsReceived
	var updateDetail = models.UpdateRating{
		Rating:                  seller.Rating,
		TotalRatings:            seller.TotalRatings,
		NumberOfRatingsReceived: seller.NumberOfRatingsReceived,
	}
	err = h.DB.UpdateSellerRating(seller.ID, &updateDetail)
	if err != nil {
		c.JSON(500, gin.H{"message": "Something went wrong, please try again"})
		c.Abort()
		return
	}
}
