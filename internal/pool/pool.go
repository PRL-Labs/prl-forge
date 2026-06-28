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

func (p *Pool) Engine() *Engine {
	return p.engine
}

func (p *Pool) OnlineWorkers() int {
	return p.workers.Count()
}
