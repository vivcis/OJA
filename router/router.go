package router

import (
	"net/http"
	"os"

	"github.com/decadevs/shoparena/handlers"
	"github.com/decadevs/shoparena/server/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ContentType string
	handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

func SetupRouter(h *handlers.Handler) (*gin.Engine, string) {
	router := gin.Default()

	apirouter := router.Group("/api/v1")

	apirouter.GET("/searchproducts", h.SearchProductHandler)
	apirouter.GET("/sellers", h.GetSellers)
	apirouter.GET("/product/:id", h.GetProductById)
	apirouter.POST("/loginbuyer", h.LoginBuyerHandler)
	apirouter.POST("/loginseller", h.LoginSellerHandler)
	apirouter.POST("/createproduct/:sellerid", h.CreateProducts)

	//All authorized routes here
	authorizedRoutesBuyer := apirouter.Group("/")
	authorizedRoutesSeller := apirouter.Group("/")
	authorizedRoutesBuyer.Use(middleware.AuthorizeBuyer(h.DB.FindBuyerByEmail, h.DB.TokenInBlacklist))
	authorizedRoutesSeller.Use(middleware.AuthorizeSeller(h.DB.FindSellerByEmail, h.DB.TokenInBlacklist))
	authorizedRoutesBuyer.PUT("/updatebuyerprofile/:id", h.UpdateBuyerProfileHandler)
	authorizedRoutesSeller.PUT("/updatesellerprofile/:id", h.UpdateSellerProfileHandler)

	//All authorized routes here
	authorizedRoutesBuyer = apirouter.Group("/")
	authorizedRoutesBuyer.Use(middleware.AuthorizeBuyer(h.DB.FindBuyerByEmail, h.DB.TokenInBlacklist))
	{
		authorizedRoutesBuyer.PUT("/updatebuyerprofile", h.UpdateBuyerProfileHandler)
		authorizedRoutesBuyer.GET("/getbuyerprofile", h.GetBuyerProfileHandler)

	}

	authorizedRoutesSeller = apirouter.Group("/")
	authorizedRoutesSeller.Use(middleware.AuthorizeSeller(h.DB.FindSellerByEmail, h.DB.TokenInBlacklist))
	{
		authorizedRoutesSeller.PUT("/updatesellerprofile", h.UpdateSellerProfileHandler)
		authorizedRoutesSeller.GET("/getsellerprofile", h.GetSellerProfileHandler)
		authorizedRoutesSeller.GET("/seller/total/product/sold", h.GetTotalSoldProductCount)
		authorizedRoutesSeller.DELETE("/deleteproduct/:id", h.DeleteSellerProduct)
	}

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8081"
	}
	return router, port
}
