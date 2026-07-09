package history

import (
	"sync"
	"time"
)

type WorkerPoint struct {
	Time     time.Time `json:"time"`
	Hashrate int64     `json:"hashrate"`
}

type WorkerManager struct {
	mu     sync.RWMutex
	points map[string][]WorkerPoint
}

func NewWorkerManager() *WorkerManager {
	return &WorkerManager{
		points: make(map[string][]WorkerPoint),
	}
}

func (m *WorkerManager) Add(wallet, worker string, hashrate int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := wallet + "." + worker

	// Не записвай еднакви стойности
	if history := m.points[key]; len(history) > 0 {
		last := history[len(history)-1]
		if last.Hashrate == hashrate {
			return
		}
	}

	m.points[key] = append(m.points[key], WorkerPoint{
		Time:     time.Now(),
		Hashrate: hashrate,
	})

	if len(m.points[key]) > 2880 {
		m.points[key] = m.points[key][1:]
	}
}

func (m *WorkerManager) Get(wallet, worker string) []WorkerPoint {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := wallet + "." + worker

	return append([]WorkerPoint(nil), m.points[key]...)
}
