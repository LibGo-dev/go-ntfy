package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func buy(ticker string, url string, auth string) {
	req, _ := http.NewRequest("POST", url,
		strings.NewReader("Buy signal received for "+ticker))
	req.Header.Set("Title", "BUY SIGNAL")
	req.Header.Set("Priority", "urgent")
	req.Header.Set("Tags", "+1,tada")
	req.Header.Set("Authorization", "Basic "+auth)
	http.DefaultClient.Do(req)
}

func sell(ticker string, url string, auth string) {
	req, _ := http.NewRequest("POST", url,
		strings.NewReader("Sell signal received for "+ticker))
	req.Header.Set("Title", "SELL SIGNAL")
	req.Header.Set("Priority", "urgent")
	req.Header.Set("Tags", "-1,no_entry")
	req.Header.Set("Authorization", "Basic "+auth)
	http.DefaultClient.Do(req)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	login := os.Getenv("NTFY_USERNAME") + ":" + os.Getenv("NTFY_PASSWORD")
	auth := base64.StdEncoding.EncodeToString([]byte(login))

	buy("$PEPE", os.Getenv("POST_URL"), auth)
	sell("$PEPE", os.Getenv("POST_URL"), auth)
}
