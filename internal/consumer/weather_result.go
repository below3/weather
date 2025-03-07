package consumer

type WeatherSummary struct {
	TempCity       string
	TempAvg        float64
	FogCity        string
	FogOccurance   int16
	ClearCity      string
	ClearOccurance int16
}

func newEmptyWeatherSummary() WeatherSummary {
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
