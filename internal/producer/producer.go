package producer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/weather/weather/internal/models"
)

// Serial Producer
func SerialProducer(filePath string, weatherChan chan<- models.WeatherModel) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		panic(err)
	}
	decoder.InputOffset()

	for decoder.More() {
		var weather models.WeatherModel
		if err := decoder.Decode(&weather); err != nil {
			fmt.Println(err)
		}
		weatherChan <- weather
	}
	close(weatherChan)
}
