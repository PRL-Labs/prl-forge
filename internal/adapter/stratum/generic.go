package stratumadapter

import "github.com/techobg/prl-forge/internal/pool"

type GenericAdapter struct{}

func NewGenericAdapter() *GenericAdapter {
	return &GenericAdapter{}
}

func (a *GenericAdapter) Notify(job *pool.Job) (any, error) {
	return NotifyMessage{
		ID:     nil,
		Method: "mining.notify",
		Params: NotifyParams{
			Header:      job.Header,
			Height:      job.Height,
			JobID:       job.ID,
			Target:      job.Target,
			CertVersion: job.CertVersion,
		},
	}, nil
}

func (a *GenericAdapter) Submit(params any) error {
	return nil
}