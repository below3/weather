package start

import (
	"encoding/json"
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
	prodNumber, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		panic(fmt.Errorf("please provide the producer number %s", err))
	}
	consNumber, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		panic(fmt.Errorf("please provide the producer number %s", err))
	}

	var wg sync.WaitGroup
	weatherSummary := consumer.NewEmptyWeatherSummary()
	var endMarker producer.EndMarker
	chanWeather := make(chan models.WeatherModel, prodNumber)

	for i := range consNumber {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go producer.ConcurentProducer(i, weatherFile, chanWeather, consNumber, &endMarker)
		}()
	}

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

	encoder := json.NewEncoder(fileToSave)
	err = encoder.Encode(weatherSummary)
	if err != nil {
		panic(err)
	}
}
