package start

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/weather/weather/internal/consumer"
	"github.com/weather/weather/internal/models"
	"github.com/weather/weather/internal/producer"
)

const ResultFile = "result.json"

func StartWeatherApp(weatherFile string) {
	argsWithoutProg := os.Args[1:]
	prodNumber, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(fmt.Errorf("please provide the producer number %s", err))
	}

	var wg sync.WaitGroup
	weatherSummary := consumer.NewEmptyWeatherSummary()

	chanWeather := make(chan models.WeatherModel, prodNumber)
	go producer.SerialProducer(weatherFile, chanWeather)

	for range prodNumber {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumer.ConcurentConsumer(chanWeather, &weatherSummary)
		}()
	}
	wg.Wait()

	fileToSave, err := os.Create(ResultFile)
	if err != nil {
		panic(err)
	}
	defer fileToSave.Close()

	resultString := fmt.Sprintf("Best average temperature is: %s, %v \nBest fog city: %s, %d \nBest clear sky: %s, %d",
		weatherSummary.TempCity, weatherSummary.TempAvg,
		weatherSummary.FogCity, weatherSummary.FogOccurance,
		weatherSummary.ClearCity, weatherSummary.ClearOccurance)
	fileToSave.WriteString(resultString)
}
