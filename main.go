package main

import "temperature-monitoring-go/pkg/caller"

func main() {
	caller.GetWeatherData("Lisbon")
}