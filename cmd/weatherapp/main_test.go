package main

import (
	"os"
	"testing"

	"github.com/weather/weather/internal/start"
)

const testWeather = "testData/pl172.json"

func BenchmarkEntireApp(b *testing.B) {
	os.Args = []string{"s", "4", "2"}
	for b.Loop() {
		start.StartWeatherApp(testWeather)
	}
	_ = os.Remove(start.ResultFile)
}
