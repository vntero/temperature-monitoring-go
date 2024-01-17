package publisher

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func getWeatherData() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")

	apiUrl, err := url.Parse("http://api.weatherapi.com/v1/current.json")
	if err != nil {
		fmt.Println("Error parsing API URL", err)
		return
	}

	apiUrl.Query().Add("key", apiKey)
	apiUrl.Query().Add("q", "Lisbon")

	// make a get request to the api
	response, err := http.Get(apiUrl.String())
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// check if request was successful
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Non-OK status code received -", response.Status)
		return
	}
}