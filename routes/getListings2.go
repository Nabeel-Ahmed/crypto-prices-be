package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabeel-ahmed/crypto-prices-be/services"
)

// "github.com/nabeel-ahmed/crypto-prices-be/services"

func getListings2(c *gin.Context) {
	services.CoinMarketCapService(c, http.MethodGet, "v1/cryptocurrency/listings/latest")	
}