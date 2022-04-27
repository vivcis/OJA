package handlers

import (
	"github.com/decadevs/multivendor/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerWelcomeStatement(c *gin.Context) {

	data := &models.User{
		Username: "Oluwadurotimi",
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Message": data,
	})
}
