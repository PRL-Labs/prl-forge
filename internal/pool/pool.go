package pool

import (
    workers "github.com/techobg/prl-forge/internal/updater/workers"
)

type Pool struct {
	workers *workers.Manager
	engine  *Engine
}

func New() *Pool {
	return &Pool{
		workers: workers.NewManager(),
		engine:  NewEngine(),
	}
}

func (p *Pool) Workers() *workers.Manager {
	return p.workers
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