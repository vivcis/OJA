package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/decadevs/shoparena/database"
	servererrors "github.com/decadevs/shoparena/handlers/serverErrors"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/router"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// decode decodes the body of c into v
func (s *Server) decode(c *gin.Context, v interface{}) []string {
	if err := c.ShouldBindJSON(v); err != nil {
		errs := []string{}
		verr, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range verr {
				errs = append(errs, servererrors.NewFieldError(fieldErr).String())
			}
		} else {
			errs = append(errs, "internal server error")
		}
		return errs
	}
	return nil
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

// Server serves requests to DB with router
type Server struct {
	dB     database.DB
	Router *router.Router
}

func(s *Server) handleUpdateUserDetails()  gin.HandlerFunc{
	return func(c *gin.Context) {
  if user1, exists := c.Get("user"); exists{
	  if user, ok := user1.(*models.User); ok{
		  username, email := user.Username, user.Email
		  if errs := s.decode(c, user); errs != nil{
			JSON(c, "", http.StatusBadRequest, nil, errs)
			return
		  }

		  user.Username, user.Email = username, email
		  user.Model.UpdatedAt = time.Now()
		  if err := s.dB.UpdateUser(user); err != nil{
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
}
