package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB database.DB
}

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func JSON(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message": message,
		"data":    data,
		"errors":  errs,
		"status":  http.StatusText(status),
	}

	c.JSON(status, responsedata)
}

func (h *Handler) UpdateProfileHandler(c *gin.Context){

		if user1, exists := c.Get("user"); exists {
			if user, ok := user1.(*models.User); ok {
				username, email := user.Username, user.Email
				if errs := c.BindJSON(user); errs != nil {
					JSON(c, "", http.StatusBadRequest, nil, []string{"errs"})
					return
				}

				user.Username, user.Email = username, email
				user.UpdatedAt = time.Now()
				if err := h.DB.UpdateUser(user); err != nil {
					log.Printf("update user error : %v\n", err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
					return
				}
				JSON(c, "user updated successfuly", http.StatusOK, nil, nil)
				return
			}
		}
		log.Printf("can't get user from context\n")
		JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
	}


func (h *Handler) SearchProductHandler(c *gin.Context) {
	//Equivalent to param
	category := c.Query("category")
	lowerPrice := c.Query("lower-price")
	upperPrice := c.Query("upper-price")
	name := c.Query("name")

	product, err := h.DB.SearchProduct(lowerPrice, upperPrice, category, name)
	if err != nil {
		log.Println("handler error in search product", err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(product) == 0 {
		c.JSON(http.StatusInternalServerError, "no such product")
		return
	}

	c.JSON(http.StatusFound, product)
}

