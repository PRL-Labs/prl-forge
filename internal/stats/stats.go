package stats

import (
	"sync"

	"github.com/techobg/prl-forge/internal/pearl"
)

var (
	mu       sync.RWMutex
	template *pearl.BlockTemplate
)

func Update(t *pearl.BlockTemplate) {
	mu.Lock()
	defer mu.Unlock()

	template = t
}

func Template() *pearl.BlockTemplate {
	mu.RLock()
	defer mu.RUnlock()

	return template
}