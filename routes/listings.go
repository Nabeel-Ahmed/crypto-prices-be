package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getListings(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	coinGeckoAPIKey := os.Getenv("COIN_GECKO_API_KEY")

	println(coinGeckoAPIKey)

	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(coinGeckoAPIKey, "should be here")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinGeckoAPIKey)

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
