package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)


func (h *Handler) UpdateSellerProfileHandler(c *gin.Context) {
	if userI, exists := c.Get("user"); exists {
		user, ok := userI.(*models.Seller)
		if  ok {
			var update models.UpdateUser
			email := user.Email
			
			if errs := h.decode(c, &update); errs != nil {
				c.JSON(http.StatusBadRequest, errs)
				return
			}
			user.Email = email
			user.UpdatedAt = time.Now()
			if err := h.DB.UpdateSellerProfile(user.ID, &update); err != nil {
				log.Printf("update seller error : %v\n", err)
				c.JSON(http.StatusInternalServerError, []string{"update seller not successful"})
				return
			}
			c.JSON(http.StatusOK, "seller updated successfuly")
			return
		}
	
		log.Printf("can't get seller from context\n")
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
	}
	
}
