package pool

import (
	"log"
	"math/big"
	"sync"
	"time"
)

type RoundStats struct {
	mu sync.RWMutex

	shares uint64
	work   *big.Int

	started    time.Time
	lastSample time.Time
	lastWork   *big.Int
}

func NewRoundStats() *RoundStats {
	now := time.Now()

	return &RoundStats{
		work:       new(big.Int),
		lastWork:   new(big.Int),
		started:    now,
		lastSample: now,
	}
}

func (r *RoundStats) AddWork(work *big.Int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares++

	r.work.Add(r.work, work)

	log.Printf(
		"ADD WORK -> shares=%d work=%s",
		r.shares,
		r.work.String(),
	)
}

func (r *RoundStats) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares = 0

	r.work = new(big.Int)
	r.lastWork = new(big.Int)

	now := time.Now()

	r.started = now
	r.lastSample = now
}

func (r *RoundStats) Shares() uint64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.shares
}

func (r *RoundStats) Work() *big.Int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return new(big.Int).Set(r.work)
}

func (r *RoundStats) SampleWork() (*big.Int, float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(r.lastSample).Seconds()
	if elapsed <= 0 {
		return new(big.Int), 0
	}

	work := new(big.Int).Sub(r.work, r.lastWork)

	r.lastWork.Set(r.work)
	r.lastSample = now

	return work, elapsed
}

func (r *RoundStats) Luck(networkDifficulty float64) float64 {
	return 0
}