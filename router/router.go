package router

import (
	"net/http"
	"os"

	"github.com/decadevs/shoparena/updatepasswordhandlers"
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
	apirouter.PUT("/buyer/resetpassword/:email", h.BuyerResetPassword)
	apirouter.PUT("/seller/resetpassword/:email", h.SellerResetPassword)
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
