package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalAvgTempSuccess(t *testing.T) {
	w := WeatherModel{
		Temperature: []float64{1, 2, 3, 4, 5},
	}
	result := w.CalAvgTemp()
	assert.Equal(t, float64(3), result)
}

func TestCheckWeaterCodesSuccess(t *testing.T) {
	w := WeatherModel{
		WeatherCode: []int16{1, 2, 3, 4, 5, 0, 0, 0, 0, 45, 45, 1, 2, 3, 45, 0},
	}
	fog, clear := w.CheckWeaterCodes()
	assert.Equal(t, int16(3), fog, "Inccorect Fog Count")
	assert.Equal(t, int16(5), clear, "Inccorect Clear sky Count")
}
