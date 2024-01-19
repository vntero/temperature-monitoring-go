package main

import "temperature-monitoring-go/pkg/publisher"

func main() {
	publisher.GetWeatherData("Lisbon")
}