package start

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weather/weather/internal/consumer"
)

func TestStartWeatherApp(t *testing.T) {
	os.Args = []string{"s", "2", "2"}
	StartWeatherApp("testData/2Entry.json")
	file, err := os.Open(ResultFile)
	assert.NoError(t, err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	var weather consumer.WeatherSummary
	err = decoder.Decode(&weather)
	assert.NoError(t, err)

	weatherExp := consumer.WeatherSummary{
		TempCity:       "Gliwice",
		TempAvg:        9.3,
		FogCity:        "Gliwice",
		FogOccurance:   1,
		ClearCity:      "None",
		ClearOccurance: 0,
	}

	assert.Equal(t, weatherExp, weather)

	err = os.Remove(ResultFile)
	assert.NoError(t, err)

}
