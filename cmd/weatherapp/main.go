package main

import (
	"github.com/weather/weather/internal/starter"
)

const weatherFile = "startFiles/pl172.json"

func main() {
	starter.StartWeatherApp(weatherFile)
}
