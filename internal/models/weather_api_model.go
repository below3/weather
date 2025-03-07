package models

type WeatherModelFromApi struct {
	Hourly struct {
		Temperature []float64 `json:"temperature_2m"`
		WeatherCode []int16   `json:"weather_code"`
	}
}
