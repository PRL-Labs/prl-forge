package pool

import (
	"sync"
	"time"
)

type ShareWindow struct {
	mu     sync.RWMutex
	shares []time.Time
}

func NewShareWindow() *ShareWindow {
	return &ShareWindow{
		shares: make([]time.Time, 0, 10000),
	}
}