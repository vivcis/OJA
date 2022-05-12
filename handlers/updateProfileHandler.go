package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) UpdateProfileHandler(c *gin.Context){
	details := &struct{
       Name string `json:"user_name"`
	   Email string `json:"email"`
	   NewEmail string `json:"new_email"`

	}{}
				if errs := c.BindJSON(details); errs != nil {
					c.JSON( http.StatusBadRequest,[]string{"errs"})
					return
				}
				username := details.Name
				buyerDetail, err := h.DB.FindBuyerByEmail(details.Email)
				if err != nil{
					log.Println("buyer not found", err)
					c.JSON(http.StatusBadRequest,gin.H{
					"error":[]string{"internal server error"},
					})
					return
				}
				log.Print(buyerDetail.User)
				buyerDetail.User.Username = username
				buyerDetail.User.Email = details.NewEmail
				if err := h.DB.UpdateUser(buyerDetail, details.Email); err != nil {
					log.Printf("update user error : %v\n", err)
					c.JSON ("", http.StatusInternalServerError, nil, []string{"update user error"})
					return
				}
				c.JSON("user updated successfuly", http.StatusOK, nil, nil)
				
		// if user1, exists := c.Get("user"); exists {
		// }
		
	}

