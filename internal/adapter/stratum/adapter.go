package stratumadapter

import "github.com/techobg/prl-forge/internal/pool"

type Adapter interface {
	Notify(job *pool.Job) (any, error)
	Submit(params any) error
}