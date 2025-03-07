package main

import (
	"encoding/json"
	"os"

	"github.com/weather/weather/internal/models"
)

const weatherURL = "https://historical-forecast-api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&start_date=2024-10-06&" +
	"end_date=2025-03-06&hourly=temperature_2m,weather_code&"
const pullFile = "startFiles/pl-pull.json"

func main() {
	file, err := os.Open(pullFile)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cities []models.Cities = make([]models.Cities, 172)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cities)
	if err != nil {
		panic(err)
	}

	chanCities := make(chan models.Cities, 20)
	chanWeather := make(chan models.WeatherModel, 20)

	var weatherList []models.WeatherModel = make([]models.WeatherModel, 172)

	// Create and start 20 workers
	for i := 0; i <= 19; i++ {
		go requestWeatherData(chanCities, chanWeather)
		chanCities <- cities[i]
	}

	// Receive weather and send more cities to be queried
	for i := 0; i <= 151; i++ {
		weatherList[i] = <-chanWeather
		chanCities <- cities[i+20]
	}
	close(chanCities)

	// Receive remaining weather
	for i := 152; i <= 171; i++ {
		weatherList[i] = <-chanWeather
	}

	// Save data into a file
	fileToSave, err := os.Create("startFiles/pl172.json")
	if err != nil {
		panic(err)
	}
	defer fileToSave.Close()
	encoder := json.NewEncoder(fileToSave)
	err = encoder.Encode(weatherList)
	if err != nil {
		panic(err)
	}
}

// Will create a worker for cities channel
func requestWeatherData(reciveCities <-chan models.Cities, returnWeather chan<- models.WeatherModel) {
	for city := range reciveCities {
		weatherModel, err := city.DownloadWether(weatherURL)
		if err != nil {
			panic(err)
		}
		returnWeather <- weatherModel

	}
}
