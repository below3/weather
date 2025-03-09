package main

import (
	"os"
	"testing"

	"github.com/weather/weather/internal/start"
)

const testWeather = "testData/pl172.json"

func BenchmarkEntireApp(b *testing.B) {
	os.Args = []string{"s", "5"}
	for i := 0; i < b.N; i++ {
		start.StartWeatherApp(testWeather)
	}
	_ = os.Remove(start.ResultFile)
}
