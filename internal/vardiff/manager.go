package vardiff

import (
	"sync"
	"time"
  "log"
)

const (
	windowSize = 8
)

type Manager struct {
	mu sync.RWMutex

	TargetShareTime time.Duration

	MinDifficulty float64
	MaxDifficulty float64

	shareTimes map[string][]time.Time
	difficulty map[string]float64
}

func New() *Manager {
	return &Manager{
		TargetShareTime: 15 * time.Second,

		MinDifficulty: 1e9,
		MaxDifficulty: 1e15,

		shareTimes: make(map[string][]time.Time),
		difficulty: make(map[string]float64),
	}
}

func (m *Manager) ObserveShare(wallet string) float64 {


	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()

	diff := m.difficulty[wallet]
	if diff == 0 {
		diff = m.MinDifficulty
	}

	times := append(m.shareTimes[wallet], now)
  
  log.Printf(
    "VARDIFF WINDOW wallet=%s shares=%d",
    wallet,
    len(times),
)

	if len(times) > windowSize {
		times = times[len(times)-windowSize:]
	}

	m.shareTimes[wallet] = times

	
	if len(times) < windowSize {
		m.difficulty[wallet] = diff
		return diff
	}

	var total time.Duration

	for i := 1; i < len(times); i++ {
		total += times[i].Sub(times[i-1])
	}

	avg := total / time.Duration(len(times)-1)
  
  log.Printf(
    "VARDIFF AVG wallet=%s avg=%s",
    wallet,
    avg,
)
  

	if avg >= 12*time.Second && avg <= 18*time.Second {
		m.difficulty[wallet] = diff
		return diff
	}

	
	if avg < 12*time.Second {
		diff *= 2
	}


	if avg > 18*time.Second {
		diff /= 2
	}

	if diff < m.MinDifficulty {
		diff = m.MinDifficulty
	}

	if diff > m.MaxDifficulty {
		diff = m.MaxDifficulty
	}

	m.difficulty[wallet] = diff

	return diff
}

func (m *Manager) CurrentDifficulty(wallet string) float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	diff := m.difficulty[wallet]
	if diff == 0 {
		return m.MinDifficulty
	}

	return diff
}

func ToDisplayDifficulty(diff float64) uint64 {
	switch uint64(diff) {

	case 1_000_000_000:
		return 1 << 33 // 8G

	case 2_000_000_000:
		return 1 << 34 // 16G

	case 4_000_000_000:
		return 1 << 35 // 32G

	case 8_000_000_000:
		return 1 << 36 // 64G

	case 16_000_000_000:
		return 1 << 37 // 128G

	case 32_000_000_000:
		return 1 << 38 // 256G

	case 64_000_000_000:
		return 1 << 39 // 512G

	case 128_000_000_000:
		return 1 << 40 // 1T

	default:
		return 1 << 33
	}
}