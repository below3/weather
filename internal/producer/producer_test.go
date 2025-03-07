package producer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weather/weather/internal/models"
)

var weatherExample1 = models.WeatherModel{
	City:        "Toruń",
	Temperature: []float64{9.4, 9.2, 8.9},
	WeatherCode: []int16{3, 3, 3, 3, 3, 3},
}

var weatherExample2 = models.WeatherModel{
	City:        "Gliwice",
	Temperature: []float64{9.3, 9.3, 9.3},
	WeatherCode: []int16{3, 3, 3, 45},
}

func TestSerialProducer(t *testing.T) {
	chanWeather := make(chan models.WeatherModel, 2)
	SerialProducer("testData/2Entry.json", chanWeather)
	var weathers []models.WeatherModel
	for w := range chanWeather {
		weathers = append(weathers, w)
	}

	assert.Equal(t, weatherExample1, weathers[0])
	assert.Equal(t, weatherExample2, weathers[1])

}
