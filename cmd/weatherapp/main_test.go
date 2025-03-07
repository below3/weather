package main

import (
	"os"
	"testing"

	"github.com/weather/weather/internal/starter"
)

const testWeather = "testData/pl172.json"

func BenchmarkEntireApp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		starter.StartWeatherApp(testWeather)
	}
	_ = os.Remove("result.json")
}
