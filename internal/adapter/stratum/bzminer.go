package stratumadapter

import "github.com/techobg/prl-forge/internal/pool"

type BzMinerAdapter struct{}

func NewBzMinerAdapter() *BzMinerAdapter {
	return &BzMinerAdapter{}
}

func (a *BzMinerAdapter) Notify(job *pool.Job) (any, error) {
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

func (a *BzMinerAdapter) Submit(params any) error {
	// TODO
	return nil
}
