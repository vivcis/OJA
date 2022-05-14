package handlers

import (
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)


func (h *Handler) UpdateBuyerProfileHandler(c *gin.Context) {
	if userI, exists := c.Get("user"); exists {
		user, ok := userI.(*models.Buyer)
		if  ok {
			var update models.UpdateUser
			email := user.Email
			log.Println(user)
			
			if errs := h.decode(c, &update); errs != nil {
				c.JSON(http.StatusBadRequest, errs)
				return
			}
	
			user.Email = email
			user.UpdatedAt = time.Now()
			if err := h.DB.UpdateBuyerProfile(user.ID, &update); err != nil {
				log.Printf("update buyer error : %v\n", err)
				c.JSON(http.StatusInternalServerError, []string{"update buyer not successful"})
				return
			}
			c.JSON(http.StatusOK, "buyer updated successfuly")
			return
		}
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
	}
	
}

