package producer

import (
	"sync"
)

type EndMarker struct {
	mu    sync.Mutex
	Count int
}

func (e *EndMarker) checkAndIncrease(consCount int) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.Count != consCount-1 {
		e.Count++
		return false
	}
	return true
}
