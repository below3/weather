package consumer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weather/weather/internal/models"
)

func Test_checkSummary(t *testing.T) {
	weather1 := models.WeatherModel{
		City:        "Byd",
		Temperature: []float64{1, 2, 3, 4, 5},
		WeatherCode: []int16{0, 0, 0, 1, 2, 3},
	}
	weather2 := models.WeatherModel{
		City:        "Gru",
		Temperature: []float64{0, 0, 0, 0},
		WeatherCode: []int16{0, 0, 45},
	}

	wSum := NewEmptyWeatherSummary()

	wSum.checkSummary(weather1)
	wSum.checkSummary(weather2)

	wSumMock := WeatherSummary{
		TempCity:       weather1.City,
		TempAvg:        3,
		FogCity:        weather2.City,
		FogOccurance:   1,
		ClearCity:      weather1.City,
		ClearOccurance: 3,
	}
	assert.Equal(t, wSumMock, wSum)

}
