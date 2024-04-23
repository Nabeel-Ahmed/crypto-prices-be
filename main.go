package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type CryptoResponse struct {
	Data []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"data"`
}

func main() {
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
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
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
