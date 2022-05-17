package router

import (
	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
<<<<<<< HEAD
	apirouter.POST("buyer/forgotpassword", h.SendForgotPasswordEMailHandler)
	apirouter.PUT("buyer/forgotpassword-reset", h.ForgotPasswordResetHandler)
=======
	apirouter.GET("/sellers", h.GetSellers)
	apirouter.GET("/product/:id", h.GetProductById)
	apirouter.PUT("/buyer/resetpassword/:email", h.BuyerResetPassword)
	apirouter.PUT("/seller/resetpassword/:email", h.SellerResetPassword)
>>>>>>> 955b0e9e8cf069180b51dd21846c0ac8278066f7
	apirouter.POST("/buyersignup", h.BuyerSignUpHandler)
	apirouter.POST("/sellersignup", h.SellerSignUpHandler)
	apirouter.POST("/loginbuyer", h.LoginBuyerHandler)
	apirouter.POST("/loginseller", h.LoginSellerHandler)

	//All authorized routes here
	authorizedRoutesBuyer := apirouter.Group("/")
	authorizedRoutesBuyer.Use(middleware.AuthorizeBuyer(h.DB.FindBuyerByEmail, h.DB.TokenInBlacklist))
<<<<<<< HEAD
	authorizedRoutesSeller.Use(middleware.AuthorizeSeller(h.DB.FindSellerByEmail, h.DB.TokenInBlacklist))
	authorizedRoutesBuyer.PUT("/buyer/resetpassword/:email", h.BuyerUpdatePassword)
	authorizedRoutesSeller.PUT("/seller/resetpassword/:email", h.SellerUpdatePassword)
	authorizedRoutesBuyer.PUT("/updatebuyerprofile/:id", h.UpdateBuyerProfileHandler)
	authorizedRoutesSeller.PUT("/updatesellerprofile/:id", h.UpdateSellerProfileHandler)
=======
	{
		authorizedRoutesBuyer.PUT("/updatebuyerprofile", h.UpdateBuyerProfileHandler)
		authorizedRoutesBuyer.GET("/getbuyerprofile", h.GetBuyerProfileHandler)
		
	}
>>>>>>> 955b0e9e8cf069180b51dd21846c0ac8278066f7

	authorizedRoutesSeller := apirouter.Group("/")
	authorizedRoutesSeller.Use(middleware.AuthorizeSeller(h.DB.FindSellerByEmail, h.DB.TokenInBlacklist))
	{
		authorizedRoutesSeller.PUT("/updatesellerprofile", h.UpdateSellerProfileHandler)
		authorizedRoutesSeller.GET("/getsellerprofile", h.GetSellerProfileHandler)
		
	}
	
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
