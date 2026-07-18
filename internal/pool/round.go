package pool

import (
	"log"
	"sync"
  "time"
)

type RoundStats struct {
        mu sync.RWMutex

        shares uint64
        work   float64

        started    time.Time
        lastSample time.Time
        lastWork   float64
        
        miningSeconds time.Duration
	online         bool
	sessionStarted time.Time
        
}

func NewRoundStats() *RoundStats {
        now := time.Now()

        return &RoundStats{
                started:    now,
                lastSample: now,
        }
}

func (r *RoundStats) AddShare(diff float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares++
	r.work += diff

	log.Printf("ADD SHARE -> shares=%d work=%.0f", r.shares, r.work)
}

func (r *RoundStats) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.shares = 0
	r.work = 0
  
  now := time.Now()

r.started = now
r.lastSample = now
r.lastWork = 0
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

func (r *RoundStats) SampleWork() (float64, float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	elapsed := now.Sub(r.lastSample).Seconds()
	if elapsed <= 0 {
		return 0, 0
	}

	work := r.work - r.lastWork

	r.lastWork = r.work
	r.lastSample = now

	return work, elapsed
}

func (r *RoundStats) Luck(networkDifficulty float64) float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if networkDifficulty <= 0 {
		return 0
	}

	return (r.work / networkDifficulty) * 100
}

func (r *RoundStats) StartSession() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.online {
		return
	}

	r.online = true
	r.sessionStarted = time.Now()
}

func (r *RoundStats) StopSession() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.online {
		return
	}

	r.miningSeconds += time.Since(r.sessionStarted)
	r.online = false
}

func (r *RoundStats) EffectiveMiningSeconds() float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := r.miningSeconds

	if r.online {
		total += time.Since(r.sessionStarted)
	}

	return total.Seconds()
}

func (r *RoundStats) Online() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()


	return r.online
}

func (r *RoundStats) SessionStarted() time.Time {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.sessionStarted
}