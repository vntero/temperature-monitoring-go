package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// loads the env variables to this file
	godotenv.Load()

	// grabs the value of a specific env variable
	openWeatherApiKey := os.Getenv("API_KEY")

	// prints it to the console
	fmt.Println("Api key", openWeatherApiKey)

	var apiUrl string = "https://api.openweathermap.org"
}