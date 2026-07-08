package blocks

import (
	"log"
	"runtime/debug"
	"sync"
)

type Manager struct {
	mu     sync.RWMutex
	blocks []*Block
}

func NewManager() *Manager {
	return &Manager{
		blocks: make([]*Block, 0),
	}
}

func (m *Manager) Add(b *Block) {

	log.Printf("🔥 Blocks.Add called: %+v", b)
	log.Printf("%s", debug.Stack())

	m.mu.Lock()
	defer m.mu.Unlock()

	m.blocks = append(m.blocks, b)
}

func (m *Manager) List() []*Block {
	m.mu.RLock()
	defer m.mu.RUnlock()

	log.Printf("📦 Blocks.List(): count=%d", len(m.blocks))

	for i, b := range m.blocks {
		log.Printf("📦 [%d] %+v", i, b)
	}

	list := make([]*Block, len(m.blocks))
	copy(list, m.blocks)

	return list
}

func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.blocks)
}

func (m *Manager) Last() *Block {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.blocks) == 0 {
		return nil
	}

	return m.blocks[len(m.blocks)-1]
}
