package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type CryptoResponse struct {
	Data []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"data"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	coinGeckoAPIKey := os.Getenv("COIN_GECKO_API_KEY")

	println(coinGeckoAPIKey)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", coinGeckoAPIKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	defer resp.Body.Close() // Close the response body when done

	fmt.Println(resp.Status)

	// Decode the JSON response into CryptoResponse struct
	var cryptoResp CryptoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cryptoResp); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		os.Exit(1)
	}

	// Access the decoded data
	for _, crypto := range cryptoResp.Data {
		fmt.Printf("ID: %d, Name: %s, Symbol: %s\n", crypto.ID, crypto.Name, crypto.Symbol)
	}
}
