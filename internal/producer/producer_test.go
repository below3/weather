package producer

import (
	"sync"
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

var weatherExample3 = models.WeatherModel{
	City:        "Bydgoszcz",
	Temperature: []float64{8, 6},
	WeatherCode: []int16{1, 2, 3, 45},
}

var weatherExample4 = models.WeatherModel{
	City:        "Grudziądz",
	Temperature: []float64{7, 7, 7},
	WeatherCode: []int16{0, 0, 3, 45},
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

func TestConcurentProducer_2Producers(t *testing.T) {

	var endMarker EndMarker
	chanWeather := make(chan models.WeatherModel, 2)
	var wg sync.WaitGroup
	for i := range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go ConcurentProducer(i, "testData/2Entry.json", chanWeather, 2, &endMarker)
		}()
	}

	var weathers []models.WeatherModel
	for w := range chanWeather {
		weathers = append(weathers, w)
	}

	assert.Equal(t, weatherExample1, weathers[0])
	assert.Equal(t, weatherExample2, weathers[1])

}
func TestConcurentProducer_1Producer(t *testing.T) {

	var endMarker EndMarker
	chanWeather := make(chan models.WeatherModel, 2)
	var wg sync.WaitGroup
	for i := range 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go ConcurentProducer(i, "testData/2Entry.json", chanWeather, 1, &endMarker)
		}()
	}

	var weathers []models.WeatherModel
	for w := range chanWeather {
		weathers = append(weathers, w)
	}

	assert.Equal(t, weatherExample1, weathers[0])
	assert.Equal(t, weatherExample2, weathers[1])

}

// Check if it still works when the number of readers exceed the number of records
func TestConcurentProducer_3Producer(t *testing.T) {

	var endMarker EndMarker
	chanWeather := make(chan models.WeatherModel, 2)
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go ConcurentProducer(i, "testData/2Entry.json", chanWeather, 3, &endMarker)
		}()
	}

	var weathers []models.WeatherModel
	for w := range chanWeather {
		weathers = append(weathers, w)
	}

	assert.Equal(t, 2, len(weathers))
	assert.Contains(t, weathers, weatherExample1)
	assert.Contains(t, weathers, weatherExample2)

}

// Check if it still works when the number of readers is uneven to the number of records
func TestConcurentProducer_4Entry(t *testing.T) {

	var endMarker EndMarker
	chanWeather := make(chan models.WeatherModel, 2)
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go ConcurentProducer(i, "testData/4Entry.json", chanWeather, 3, &endMarker)
		}()
	}

	var weathers []models.WeatherModel
	for w := range chanWeather {
		weathers = append(weathers, w)
	}

	assert.Equal(t, 4, len(weathers))
	assert.Contains(t, weathers, weatherExample1)
	assert.Contains(t, weathers, weatherExample2)
	assert.Contains(t, weathers, weatherExample3)
	assert.Contains(t, weathers, weatherExample4)
}
