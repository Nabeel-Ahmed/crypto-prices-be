package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nabeel-ahmed/crypto-prices-be/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
