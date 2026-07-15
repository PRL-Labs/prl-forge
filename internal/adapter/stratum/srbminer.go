package stratumadapter

import (

  "github.com/techobg/prl-forge/internal/pool"
)

type SRBMinerAdapter struct{}

func NewSRBMinerAdapter() *SRBMinerAdapter {
	return &SRBMinerAdapter{}
}

type NotifyMessage struct {
	ID     any          `json:"id"`
	Method string       `json:"method"`
	Params NotifyParams `json:"params"`
}

type NotifyParams struct {
	Header      string `json:"header"`
	Height      int64  `json:"height"`
	JobID       string `json:"job_id"`
	Target      string `json:"target"`
	CertVersion int    `json:"cert_version"`
}

func (a *SRBMinerAdapter) Notify(job *pool.Job) (any, error) {
	return NotifyMessage{
		ID:     nil,
		Method: "mining.notify",
		Params: NotifyParams{
			Header:      job.Header,
			Height:      job.Height,
			JobID: job.ID,
			Target:      job.Target,
			CertVersion: job.CertVersion,
		},
	}, nil
}

func (a *SRBMinerAdapter) Submit(params any) error {
	// TODO: Pearl submit parser
	return nil
}
