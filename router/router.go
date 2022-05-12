package router

import (
	"net/http"
	"os"

	"github.com/decadevs/shoparena/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ContentType string
	handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

func SetupRouter(h *handlers.Handler) (*gin.Engine, string) {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/ping", handlers.PingHandler)
	apirouter.GET("/searchproducts", h.SearchProductHandler)
	apirouter.PUT("/buyer/uppdatepassword/:email", h.BuyerUpdatePasswordHandler)
	apirouter.PUT("/seller/updatepassword/:email", h.SellerUpdatePasswordHandler)
	apirouter.POST("/buyer/forgotpassword", h.SendForgotPasswordEMailHandler)
	apirouter.PUT("/buyer/forgotpassword-reset?reset_token=", h.ForgotPasswordResetHandler)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
