package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	OpenWeatherApiKey := os.Getenv("API_KEY")

	fmt.Println("Api key", OpenWeatherApiKey)
}