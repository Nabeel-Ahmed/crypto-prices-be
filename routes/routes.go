package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/listings", getListings)
	server.GET("/listings2", getListings2)
	server.GET("listings/new" , getListingsNew)

	// server.POST("/signup", signup)
	// server.POST("/login", login)
}
