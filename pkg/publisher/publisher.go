package publisher

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

/*
 * Makes a call to WeatherApi
 */
func GetWeatherData(city string) {
	// load env variables to this file
	godotenv.Load()

	// assign api key to the const
	apiKey := os.Getenv("API_KEY")

	// assign api url
	apiUrl, err := url.Parse("http://api.weatherapi.com/v1/current.json")
	if err != nil {
		fmt.Println("Error parsing API URL", err)
		return
	}

	fmt.Println("🚀 - api url without query params", apiUrl)

	// add query params to the API URL
	q := apiUrl.Query()
	q.Add("key", apiKey)
	q.Add("q", city)

	// ssign the modified query parameters back to apiUrl
	apiUrl.RawQuery = q.Encode()

	fmt.Println("🚀 - full url with which the request is made", apiUrl)

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