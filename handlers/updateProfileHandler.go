package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)

// func (h *Handler) UpdateProfileHandler(c *gin.Context){
// 	details := &struct{
//        Name string `json:"user_name"`
// 	   Email string `json:"email"`
// 	   NewEmail string `json:"new_email"`

// 	}{}
// 				if errs := c.BindJSON(details); errs != nil {
// 					c.JSON( http.StatusBadRequest,[]string{"errs"})
// 					return
// 				}
// 				username := details.Name
// 				buyerDetail, err := h.DB.FindBuyerByEmail(details.Email)
// 				if err != nil{
// 					log.Println("buyer not found", err)
// 					c.JSON(http.StatusBadRequest,gin.H{
// 					"error":[]string{"internal server error"},
// 					})
// 					return
// 				}
// 				log.Print(buyerDetail.User)
// 				buyerDetail.User.Username = username
// 				buyerDetail.User.Email = details.NewEmail
// 				if err := h.DB.UpdateUserProfile(buyerDetail, details.Email); err != nil {
// 					log.Printf("update user error : %v\n", err)
// 					c.JSON ( http.StatusInternalServerError, []string{"update user error"})
// 					return
// 				}
// 				c.JSON(http.StatusOK, "user updated successfuly")

// 		// if user1, exists := c.Get("user"); exists {
// 		// }

// 	}



func (h *Handler) UpdateProfileHandler(c *gin.Context) {

		if userI, exists := c.Get("user"); exists {
			fmt.Println("got here")
			if user, ok := userI.(*models.User); ok {
				var update models.UpdateUser
				email := user.Email
				log.Println(user)
				if errs := h.decode(c, &update); errs != nil {
					fmt.Println("I GOT HERE 2")
					c.JSON(http.StatusBadRequest, errs)
					return
				}

				user.Email = email
				user.UpdatedAt = time.Now()
				if err := h.DB.UpdateBuyerProfile(user.ID, &update); err != nil {
					log.Printf("update user error : %v\n", err)
					c.JSON(http.StatusInternalServerError, []string{"internal server error"})
					return
				}
				c.JSON(http.StatusOK, "user updated successfuly")
				return
			}
		}
		log.Printf("can't get user from context\n")
		c.JSON (http.StatusInternalServerError, []string{"internal server error"})	
}
