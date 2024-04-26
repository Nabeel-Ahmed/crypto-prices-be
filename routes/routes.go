package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/listings", getListings)

	// server.POST("/signup", signup)
	// server.POST("/login", login)
}
