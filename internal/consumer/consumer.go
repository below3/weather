package consumer

import (
	"github.com/weather/weather/internal/models"
)

func SerialConsumer(weatherChan <-chan models.WeatherModel, resultChan chan<- WeatherSummary) {
	compWeather := newEmptyWeatherSummary()
	for weather := range weatherChan {
		avgTemp := weather.CalAvgTemp()
		fogOcc, clearOcc := weather.CheckWeaterCodes()

		if avgTemp > compWeather.TempAvg {
			compWeather.TempAvg = avgTemp
			compWeather.TempCity = weather.City
		}

		if fogOcc > compWeather.FogOccurance {
			compWeather.FogOccurance = fogOcc
			compWeather.FogCity = weather.City
		}

		if clearOcc > compWeather.ClearOccurance {
			compWeather.ClearOccurance = clearOcc
			compWeather.ClearCity = weather.City
		}
	}

	resultChan <- compWeather
}
