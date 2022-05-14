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
	apirouter.GET("/seller/:id", h.HandleGetSellerShopByProfileAndProduct())
	apirouter.GET("/searchproducts", h.SearchProductHandler)
	apirouter.PUT("/uploadimage", h.UploadImageHandler)
	apirouter.PUT("/buyer/resetpassword/:email", h.BuyerResetPassword)
	apirouter.PUT("/seller/resetpassword/:email", h.SellerResetPassword)
	apirouter.POST("/buyersignup", h.BuyerSignUpHandler)
	apirouter.POST("/sellersignup", h.SellerSignUpHandler)
	apirouter.POST("/loginbuyer", h.LoginBuyerHandler)
	apirouter.POST("/loginseller", h.LoginSellerHandler)

	//All authorized routes here
	authorizedRoutesBuyer := apirouter.Group("/")
	authorizedRoutesSeller := apirouter.Group("/")
	authorizedRoutesBuyer.Use(middleware.AuthorizeBuyer(h.DB.FindBuyerByEmail, h.DB.TokenInBlacklist))
	authorizedRoutesSeller.Use(middleware.AuthorizeSeller(h.DB.FindSellerByEmail, h.DB.TokenInBlacklist))
	authorizedRoutesBuyer.PUT("/updatebuyerprofile/:id", h.UpdateBuyerProfileHandler)
    authorizedRoutesSeller.PUT("/updatesellerprofile/:id", h.UpdateSellerProfileHandler)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
