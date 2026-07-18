package hashrate

import (
	"sync"
	"time"
)

type Sample struct {
	Time       time.Time
	Difficulty float64
}

type Estimator struct {
	mu      sync.Mutex
	samples map[string][]Sample
}

func New() *Estimator {
	return &Estimator{
		samples: make(map[string][]Sample),
	}
}

func (e *Estimator) Add(id string, diff float64) {
	e.mu.Lock()
	defer e.mu.Unlock()

	now := time.Now()

	s := append(e.samples[id], Sample{
		Time:       now,
		Difficulty: diff,
	})

	// keep last 10 minutes
	cutoff := now.Add(-10 * time.Minute)

	n := 0
	for _, x := range s {
		if x.Time.After(cutoff) {
			s[n] = x
			n++
		}
	}

	e.samples[id] = s[:n]
}

func (e *Estimator) Work(id string, window time.Duration) float64 {
	e.mu.Lock()
	defer e.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-window)

	work := 0.0

	for _, s := range e.samples[id] {
		if s.Time.After(cutoff) {
			work += s.Difficulty
		}
	}

	return work
}

func (e *Estimator) Hashrate(id string, window time.Duration) float64 {
	work := e.Work(id, window)

	seconds := window.Seconds()
	if seconds == 0 {
		return 0
	}

	return work * 4294967296.0 / seconds
}

