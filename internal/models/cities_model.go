package models

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

const expectedArrSize = 3648

type Cities struct {
	Latitude   string `json:"lat"`
	Longtitude string `json:"lng"`
	City       string `json:"city"`
}

func (c Cities) DownloadWether(weatherURL string) (WeatherModel, error) {
	url := fmt.Sprintf(weatherURL, c.Latitude, c.Longtitude)
	response, err := http.Get(url)
	if err != nil {
		return WeatherModel{}, err
	}
	return c.parseWeather(response)
}

func (c Cities) parseWeather(response *http.Response) (WeatherModel, error) {
	waetherFromApi := WeatherModelFromApi{}
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&waetherFromApi)
	if err != nil {
		return WeatherModel{}, err
	}

	weather := WeatherModel{
		City:        c.City,
		Temperature: make([]float64, expectedArrSize),
		WeatherCode: make([]int16, expectedArrSize),
	}
	for i := range waetherFromApi.Hourly.Temperature {
		weather.Temperature[i] = math.Round(waetherFromApi.Hourly.Temperature[i]*100) / 100
		weather.WeatherCode[i] = waetherFromApi.Hourly.WeatherCode[i]
	}
	return weather, nil
}
