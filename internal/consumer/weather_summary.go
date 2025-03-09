package consumer

import (
	"sync"

	"github.com/weather/weather/internal/models"
)

type WeatherSummary struct {
	Mu             sync.Mutex `json:"-"`
	TempCity       string     `json:"bestTempCity"`
	TempAvg        float64    `json:"bestTempAvg"`
	FogCity        string     `json:"bestFogCity"`
	FogOccurance   int16      `json:"fogOccurance"`
	ClearCity      string     `json:"bestClearCity"`
	ClearOccurance int16      `json:"clearOccurance"`
}

func NewEmptyWeatherSummary() WeatherSummary {
	none := "None"
	return WeatherSummary{
		TempCity:       none,
		TempAvg:        -300,
		FogCity:        none,
		FogOccurance:   0,
		ClearCity:      none,
		ClearOccurance: 0,
	}
}

func (w *WeatherSummary) checkSummary(weather models.WeatherModel) {
	avgTemp := weather.CalAvgTemp()
	fogOcc, clearOcc := weather.CheckWeaterCodes()

	w.Mu.Lock()
	defer w.Mu.Unlock()
	if avgTemp > w.TempAvg {
		w.TempAvg = avgTemp
		w.TempCity = weather.City
	}

	if fogOcc > w.FogOccurance {
		w.FogOccurance = fogOcc
		w.FogCity = weather.City
	}

	if clearOcc > w.ClearOccurance {
		w.ClearOccurance = clearOcc
		w.ClearCity = weather.City
	}
}
