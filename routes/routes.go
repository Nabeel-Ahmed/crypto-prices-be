package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nabeel-ahmed/crypto-prices-be/services"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/listings", getListings)
	server.GET("/listings2", getListings2)
	server.GET("/listings/new" , getListingsNew)
	server.GET("/exchange/assets", func(c *gin.Context) {
		services.CoinMarketCapService(c, "GET", "v1/exchange/assets")
	})
	server.GET("/info", func(c *gin.Context) {
		services.CoinMarketCapService(c, "GET", "v2/cryptocurrency/info/1")
	})
	// server.POST("/signup", signup)
	// server.POST("/login", login)
}
