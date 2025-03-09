package start

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartWeatherApp(t *testing.T) {
	os.Args = []string{"s", "2"}
	StartWeatherApp("testData/2Entry.json")
	data, err := os.ReadFile(ResultFile)
	assert.NoError(t, err)

	assert.Equal(t,
		"Best average temperature is: Gliwice, 9.3 \nBest fog city: Gliwice, 1 \nBest clear sky: None, 0",
		string(data))
	if err = os.Remove(ResultFile); err != nil {
		assert.NoError(t, err)
	}

}
