package handlers

import (
	"fmt"
	"github.com/decadevs/shoparena/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) LoginSellerHandler(c *gin.Context) {

	type User struct {
		Password string `json:"password" binding:"required"`
		Username string `json:"username" binding:"required"`
	}

	var newSeller User
	err := c.ShouldBindJSON(&newSeller)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	hashpass, err := HashPassword(newSeller.Password)
	if err != nil {
		//write error
		c.JSON(500, gin.H{"message": "internal Server Error"})
		return
	}
	buyer, sqlErr := h.DB.FindSellerByUsername(newSeller.Username)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	fmt.Println(buyer, "the buyer", newSeller.Username)
	fmt.Println(hashpass, "1", buyer.PasswordHash, "2")

	if err := bcrypt.CompareHashAndPassword([]byte(buyer.PasswordHash), []byte(newSeller.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid Password"})
		return
	} else {
		fmt.Println("user before", buyer.ID)
		token := services.GenerateTokenWithClaims(buyer.Username)
		c.SetCookie("token", token, 3600, "/", "", false, false)
		c.JSON(http.StatusOK, gin.H{"msg": "Successfully SignedIN", "token": token, "username": buyer.Username})

	}
}
