package models

const fogCode = 45
const clearSky = 0

type WeatherModel struct {
	City        string    `json:"city"`
	Temperature []float64 `json:"temperature"`
	WeatherCode []int16   `json:"weather_code"`
}

func (w WeatherModel) CalAvgTemp() float64 {
	var sum float64
	for _, temp := range w.Temperature {
		sum += temp
	}
	return sum / float64(len(w.Temperature))
}

func (w WeatherModel) CheckWeaterCodes() (fogOcc int16, clearOcc int16) {
	for _, code := range w.WeatherCode {
		if code == fogCode {
			fogOcc++
		} else if code == clearSky {
			clearOcc++
		}
	}
	return fogOcc, clearOcc
}
