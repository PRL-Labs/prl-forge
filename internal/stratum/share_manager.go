package stratum

import "sync"

type ShareManager struct {
	mu sync.RWMutex

	shares []*Share
}

func NewShareManager() *ShareManager {
	return &ShareManager{
		shares: make([]*Share, 0),
	}
}

func (sm *ShareManager) Add(share *Share) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.shares = append(sm.shares, share)
}

func (sm *ShareManager) Count() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return len(sm.shares)
}