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
  hashrate float64
   totalDifficulty float64
   totalDiff float64
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


func (r *RoundStats) Hashrate() float64 {
	r.mu.RLock()
	work := new(big.Int).Set(r.work)
	started := r.started
	r.mu.RUnlock()

	seconds := time.Since(started).Seconds()
	if seconds <= 0 {
		return 0
	}

	// work -> float
	workFloat := new(big.Float).SetInt(work)

	// work > hashes (divide by 2^32)
	hashes := new(big.Float).Quo(
		workFloat,
		new(big.Float).SetFloat64(4294967296),
	)

	// hashes / seconds
	hashrateFloat := new(big.Float).Quo(
		hashes,
		big.NewFloat(seconds),
	)

	newHashrate, _ := hashrateFloat.Float64()

	// smoothing
	alpha := 0.3
	r.hashrate = r.hashrate*(1-alpha) + newHashrate*alpha

	return r.hashrate
}
func (r *RoundStats) Luck(networkDifficulty float64) float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if networkDifficulty <= 0 {
		return 0
	}

	// expected shares = difficulty (diff1 shares)
	expectedShares := networkDifficulty

	if expectedShares == 0 {
		return 0
	}

	luck := float64(r.shares) / expectedShares

	return luck * 100
}

func (r *RoundStats) AddDifficulty(diff float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.totalDiff += diff
}