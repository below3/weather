package producer

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitJson(t *testing.T) {
	testString := "[{first},{second},{third}]\000"
	b := strings.NewReader(testString)
	s := bufio.NewScanner(b)
	s.Split(splitForJson())
	s.Scan()
	assert.Equal(t, "[{first}", s.Text())
	s.Scan()
	assert.Equal(t, "{second}", s.Text())
	s.Scan()
	assert.Equal(t, "{third}", s.Text())
	s.Scan()
	assert.Equal(t, "", s.Text())
}
