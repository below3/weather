package main

import (
	"github.com/weather/weather/internal/start"
)

const weatherFile = "startFiles/pl172.json"

func main() {
	start.StartWeatherApp(weatherFile)
}
