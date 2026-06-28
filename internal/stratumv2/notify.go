package stratumv2

import (
	"encoding/json"
	"log"
)

type NotifyMessage struct {
	ID     any    `json:"id"`
	Method string `json:"method"`
	Params Params `json:"params"`
}

type Params struct {
	Header      string `json:"header"`
	Height      int64  `json:"height"`
	JobID       string `json:"job_id"`
	Target      string `json:"target"`
	CertVersion int    `json:"cert_version"`
}

func NotifyJob(s *Session, job *Job) error {
	msg := NotifyMessage{
		ID:     nil,
		Method: "mining.notify",
		Params: Params{
			Header:      job.Header,
			Height:      job.Height,
			JobID:       job.ID,
			Target:      job.Target,
			CertVersion: job.CertVersion,
		},
	}

	b, _ := json.MarshalIndent(msg, "", "  ")
	log.Printf("SEND >>> %s", string(b))

	return s.Send(msg)
}
