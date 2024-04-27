package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CoinMarketCapService(c *gin.Context, requestType string, endpoint string,) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	coinMarketCapAPIKey := os.Getenv("COIN_MARKET_CAP_API_KEY")

	url := fmt.Sprintf("https://sandbox-api.coinmarketcap.com/%s", endpoint)

	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinMarketCapAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var body interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, body)
}
