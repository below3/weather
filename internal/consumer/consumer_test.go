package consumer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weather/weather/internal/models"
)

var weatherExample1 = models.WeatherModel{
	City:        "Toruń",
	Temperature: []float64{9, 9, 9},
	WeatherCode: []int16{0, 0, 3, 3, 3, 3},
}

var weatherExample2 = models.WeatherModel{
	City:        "Gliwice",
	Temperature: []float64{2, 2, 2},
	WeatherCode: []int16{0, 3, 3, 45},
}

var weatherSummaryExp = WeatherSummary{
	TempCity:       "Toruń",
	TempAvg:        9,
	FogCity:        "Gliwice",
	FogOccurance:   1,
	ClearCity:      "Toruń",
	ClearOccurance: 2,
}

func TestSerialConsumer(t *testing.T) {
	chanWeather := make(chan models.WeatherModel, 2)
	chanResult := make(chan WeatherSummary)
	go SerialConsumer(chanWeather, chanResult)
	chanWeather <- weatherExample1
	chanWeather <- weatherExample2
	close(chanWeather)
	result := <-chanResult
	assert.Equal(t, weatherSummaryExp, result)
}
