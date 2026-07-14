package pool

import (
	"github.com/techobg/prl-forge/internal/history"
	blocks "github.com/techobg/prl-forge/internal/updater/blocks"
	workers "github.com/techobg/prl-forge/internal/updater/workers"
)

type Pool struct {
	workers     *workers.Manager
	blocks      *blocks.Manager
	round       *RoundStats
  minerRounds *MinerRoundManager
	engine      *Engine
	history     *history.PoolManager
	roundHeight int64
}

func New() *Pool {
	return &Pool{
		workers: workers.NewManager(),
		blocks:  blocks.NewManager(),
		round:   NewRoundStats(),
    minerRounds: NewMinerRoundManager(),
		engine:  NewEngine(),
		history: history.NewPool(),
	}
}

func (p *Pool) Workers() *workers.Manager {
	return p.workers
}

func (p *Pool) History() *history.PoolManager {
	return p.history
}

func (p *Pool) OnlineWorkers() int {
	if p == nil || p.workers == nil {
		return 0
	}

	return p.workers.Count()
}

func (p *Pool) TotalHashrate() int64 {
	if p == nil || p.workers == nil {
		return 0
	}

	var total float64

	for _, worker := range p.workers.List() {
		total += worker.Hashrate
	}

	return int64(total)
}
func (p *Pool) Blocks() *blocks.Manager {
	return p.blocks
}

func (p *Pool) Round() *RoundStats {
	return p.round
}
func (p *Pool) MinerRounds() *MinerRoundManager {
	return p.minerRounds
}

func (p *Pool) SetRoundHeight(height int64) {
	p.roundHeight = height
}

func (p *Pool) RoundHeight() int64 {
	return p.roundHeight
}
