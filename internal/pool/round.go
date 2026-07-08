package pool

import "sync"

type RoundStats struct {
	mu sync.RWMutex

	shares uint64
	work   float64
}

func NewRoundStats() *RoundStats {
	return &RoundStats{}
}

func (r *RoundStats) AddShare(diff float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares++
	r.work += diff
}

func (r *RoundStats) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares = 0
	r.work = 0
}

func (r *RoundStats) Shares() uint64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.shares
}

func (r *RoundStats) Work() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.work
}

func (r *RoundStats) Luck(networkDifficulty float64) float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if networkDifficulty <= 0 {
		return 0
	}

	return (r.work / networkDifficulty) * 100
}
