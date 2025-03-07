package models

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWeatherSuccess(t *testing.T) {
	c := Cities{
		City: "Warsaw",
	}
	w := httptest.NewRecorder()
	w.Write([]byte(`{"city":"Warsaw"}`))
	resp := w.Result()

	weather, err := c.parseWeather(resp)
	assert.NoError(t, err)

	weatherExp := WeatherModel{City: "Warsaw",
		Temperature: make([]float64, expectedArrSize),
		WeatherCode: make([]int16, expectedArrSize)}
	assert.Equal(t, weatherExp, weather)
}

func TestParseWeatherInccorectRequest(t *testing.T) {
	c := Cities{
		City: "Warsaw",
	}
	w := httptest.NewRecorder()
	w.Write([]byte(`{"WrongString"}`))
	resp := w.Result()

	_, err := c.parseWeather(resp)
	assert.Error(t, err)
}
