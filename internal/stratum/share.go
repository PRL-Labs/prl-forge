package stratum

import "time"

type Share struct {
	Worker      string
	Wallet      string
	JobID       string
	ExtraNonce2 string
	NTime       string
	Nonce       string

	Difficulty float64

	Accepted bool

	Time time.Time
}
