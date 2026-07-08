package history

import (
	"sync"
	"time"
)

type Point struct {
	Time     time.Time `json:"time"`
	Hashrate int64     `json:"hashrate"`
}

type Manager struct {
	mu     sync.RWMutex
	points map[string][]Point
}

func New() *Manager {
	return &Manager{
		points: make(map[string][]Point),
	}
}

func (m *Manager) Add(wallet string, hashrate int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.points[wallet] = append(m.points[wallet], Point{
		Time:     time.Now(),
		Hashrate: hashrate,
	})

	if len(m.points[wallet]) > 2880 {
		m.points[wallet] = m.points[wallet][1:]
	}
}

func (m *Manager) Get(wallet string) []Point {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return append([]Point(nil), m.points[wallet]...)
}
