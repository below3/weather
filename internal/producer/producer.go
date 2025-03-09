package producer

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/weather/weather/internal/models"
)

// Could use something that like.. but It's not really needed there :)
type AcceptedConsTypes interface {
	models.WeatherModel
}

// Serial Producer
func SerialProducer(filePath string, weatherChan chan<- models.WeatherModel) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if _, err := decoder.Token(); err != nil {
		panic(err)
	}

	for decoder.More() {
		var weather models.WeatherModel
		if err := decoder.Decode(&weather); err != nil {
			panic(err)
		}
		weatherChan <- weather
	}
	close(weatherChan)
}

// Concurent Producer that does not need syncronization for reading purposes,
// but relies upon spliting the file by offset and ignoring the first return for not first producer.
// If the number of records is lower then the number of producers, can cause record to be read more than one time,
// but this is not a problem in this solution.
func ConcurentProducer[consType AcceptedConsTypes](i int, filePath string, weatherChan chan<- consType, consCount int, end *EndMarker) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileStats, _ := file.Stat()
	filePart := fileStats.Size() / int64(consCount)
	file.Seek(filePart*int64(i), 0)

	s := bufio.NewScanner(file)
	s.Split(splitForJson())
	// create an initial byte with the approx
	buf := make([]byte, 30000)
	s.Buffer(buf, 40000)

	// Custom read for the first entry due to presence of [
	if i == 0 {
		var weather consType
		s.Scan()
		currScan := s.Bytes()
		currScan = currScan[1:]
		err = json.Unmarshal(currScan, &weather)
		if err != nil {
			panic(err)
		}
		weatherChan <- weather
		filePart -= int64(len(s.Bytes())) - 1
	}

	var readBytes int64
	// Discard the first read as it will be somewhere in the json
	// But make sure the previous worker reads one more line
	if i != 0 {
		s.Scan()
		filePart -= int64(len(s.Bytes())) + 1
	}

	for readBytes < filePart {
		s.Scan()
		var weather consType
		currScan := s.Bytes()
		lenBytes := int64(len(currScan))
		readBytes += lenBytes + 1

		// dicard empty, this aroses due to modification of the read when we encounter EOF
		if lenBytes == 0 {
			break
		}

		err := json.Unmarshal(currScan, &weather)
		if err != nil {
			panic(err)
		}
		weatherChan <- weather
	}

	// Check if we can close the channel
	if end.checkAndIncrease(consCount) {
		close(weatherChan)
	}
}
