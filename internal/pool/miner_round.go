package pool

import (
	"encoding/json"
	"os"
	"sync"
)

type MinerRoundData struct {
	Shares uint64  `json:"shares"`
	Work   float64 `json:"work"`
}

type MinerRoundManager struct {
	mu     sync.RWMutex
	rounds map[string]*RoundStats
}

func NewMinerRoundManager() *MinerRoundManager {
	return &MinerRoundManager{
		rounds: make(map[string]*RoundStats),
	}
}

func (m *MinerRoundManager) AddShare(wallet string, diff float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	round, ok := m.rounds[wallet]
	if !ok {
		round = NewRoundStats()
		m.rounds[wallet] = round
	}

	round.AddShare(diff)
}

func (m *MinerRoundManager) Get(wallet string) *RoundStats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.rounds[wallet]
}

func (m *MinerRoundManager) Reset(wallet string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.rounds, wallet)
}

func (m *MinerRoundManager) Snapshot() map[string]*RoundStats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make(map[string]*RoundStats, len(m.rounds))

	for wallet, round := range m.rounds {
		copy := *round
		out[wallet] = &copy
	}

	return out
}

func (m *MinerRoundManager) Export() map[string]MinerRoundData {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make(map[string]MinerRoundData, len(m.rounds))

	for wallet, round := range m.rounds {
		out[wallet] = MinerRoundData{
			Shares: round.Shares(),
			Work:   round.Work(),
		}
	}

	return out
}

func (m *MinerRoundManager) Save(path string) error {
	data, err := json.MarshalIndent(m.Export(), "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

