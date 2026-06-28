package workers

import "sync"

type Manager struct {
	mu      sync.RWMutex
	workers map[string]*Worker
}

func NewManager() *Manager {
	return &Manager{
		workers: make(map[string]*Worker),
	}
}

func (m *Manager) Add(worker *Worker) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.workers[worker.ID] = worker
}

func (m *Manager) Remove(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.workers, id)
}

func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.workers)
}
