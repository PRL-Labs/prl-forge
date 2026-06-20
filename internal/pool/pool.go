package pool

import (
	"github.com/techobg/prl-forge/internal/workers"
)

type Pool struct {
	workers *workers.Manager
}

func New() *Pool {
	return &Pool{
		workers: workers.NewManager(),
	}
}

func (p *Pool) Workers() *workers.Manager {
	return p.workers
}

func (p *Pool) OnlineWorkers() int {
	return p.workers.Count()
}