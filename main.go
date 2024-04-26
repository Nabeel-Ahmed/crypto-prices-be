package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nabeel-ahmed/crypto-prices-be/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
