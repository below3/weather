package starter

import (
	"fmt"
	"os"

	"github.com/weather/weather/internal/consumer"
	"github.com/weather/weather/internal/models"
	"github.com/weather/weather/internal/producer"
)

const fileName = "result.json"

func StartWeatherApp(weatherFile string) {
	chanWeather := make(chan models.WeatherModel, 2)
	chanResult := make(chan consumer.WeatherSummary)
	go producer.SerialProducer(weatherFile, chanWeather)
	go consumer.SerialConsumer(chanWeather, chanResult)
	w := <-chanResult

	fileToSave, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fileToSave.Close()

	resultString := fmt.Sprintf("Best average temperature is: %s, %v \nBest fog city: %s, %d \nBest clear sky: %s, %d",
		w.TempCity, w.TempAvg,
		w.FogCity, w.FogOccurance,
		w.ClearCity, w.ClearOccurance)
	fileToSave.WriteString(resultString)
}
