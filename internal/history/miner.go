package history

import (
	"sync"
	"time"
  "log"
)

type MinerPoint struct {
	Time     time.Time `json:"time"`
	Hashrate int64     `json:"hashrate"`
}

type MinerManager struct {
	mu     sync.RWMutex
	points map[string][]MinerPoint
}

func NewMinerManager() *MinerManager {
	return &MinerManager{
		points: make(map[string][]MinerPoint),
	}
}



func (m *MinerManager) Add(wallet string, hashrate int64) {

log.Printf("MINER HISTORY ADD -> wallet=%s hashrate=%d", wallet, hashrate)

	m.mu.Lock()
	defer m.mu.Unlock()

	if history := m.points[wallet]; len(history) > 0 {
		last := history[len(history)-1]
		if last.Hashrate == hashrate {
			return
		}
	}

	m.points[wallet] = append(m.points[wallet], MinerPoint{
		Time:     time.Now(),
		Hashrate: hashrate,
	})

	if len(m.points[wallet]) > 2880 {
		m.points[wallet] = m.points[wallet][1:]
	}
}

func (m *MinerManager) Get(wallet string) []MinerPoint {

log.Printf("MINER HISTORY GET -> wallet=%s points=%d", wallet, len(m.points[wallet]))


	m.mu.RLock()
	defer m.mu.RUnlock()

	return append([]MinerPoint(nil), m.points[wallet]...)
}
