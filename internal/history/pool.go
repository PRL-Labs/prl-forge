package history

import (
	"sync"
	"time"
)

type PoolPoint struct {
	Time     time.Time `json:"time"`
	Hashrate int64     `json:"hashrate"`
}

type PoolManager struct {
	mu     sync.RWMutex
	points []PoolPoint
}

func NewPool() *PoolManager {
	return &PoolManager{
		points: make([]PoolPoint, 0),
	}
}

func (m *PoolManager) Add(hashrate int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.points) > 0 {
		last := m.points[len(m.points)-1]

		if last.Hashrate == hashrate {
			return
		}
	}

	m.points = append(m.points, PoolPoint{
		Time:     time.Now(),
		Hashrate: hashrate,
	})

	// Keep last 24h (2880 samples @ 30 sec)
	if len(m.points) > 2880 {
		m.points = m.points[1:]
	}
}

func (m *PoolManager) Get() []PoolPoint {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return append([]PoolPoint(nil), m.points...)
}
