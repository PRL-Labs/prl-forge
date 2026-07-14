package history

import (
	"sync"
	"time"
)

type ActivityPoint struct {
	Time       time.Time `json:"time"`
	Difficulty float64   `json:"difficulty"`
}

type ActivityValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

type ActivityManager struct {
	mu      sync.RWMutex
	points  map[string][]ActivityPoint
	counter map[string]float64
}

func NewActivityManager() *ActivityManager {
	return &ActivityManager{
		points:  make(map[string][]ActivityPoint),
		counter: make(map[string]float64),
	}
}

func (a *ActivityManager) Add(wallet string, difficulty float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.counter[wallet] += difficulty
}

func (a *ActivityManager) Flush() {
	a.mu.Lock()
	defer a.mu.Unlock()

	now := time.Now()

	for wallet, value := range a.counter {

		a.points[wallet] = append(a.points[wallet], ActivityPoint{
			Time:       now,
			Difficulty: value,
		})

		if len(a.points[wallet]) > 1000 {
			a.points[wallet] = a.points[wallet][1:]
		}

		a.counter[wallet] = 0
	}
}

func (a *ActivityManager) Get(wallet string) []ActivityValue {
	a.mu.RLock()
	defer a.mu.RUnlock()

	history := a.points[wallet]
  
  max := 1.0

for _, p := range history {
	if p.Difficulty > max {
		max = p.Difficulty
	}
}

	result := make([]ActivityValue, 0, len(history)*6)

	for _, p := range history {
  
  peak := (p.Difficulty / max) * 100

		

		result = append(result, ActivityValue{
			Time:  p.Time,
			Value: peak,
		})

		result = append(result, ActivityValue{
			Time:  p.Time.Add(1 * time.Second),
			Value: peak * 0.60,
		})

		result = append(result, ActivityValue{
			Time:  p.Time.Add(2 * time.Second),
			Value: peak * 0.35,
		})

		result = append(result, ActivityValue{
			Time:  p.Time.Add(3 * time.Second),
			Value: peak * 0.15,
		})

		result = append(result, ActivityValue{
			Time:  p.Time.Add(4 * time.Second),
			Value: peak * 0.05,
		})
	}

	return result
}