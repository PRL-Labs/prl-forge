package workers

import "time"

type Worker struct {
	ID     string
	Wallet string
	Name   string
	IP     string

	Hashrate float64
	Shares   uint64

	ConnectedAt time.Time
	LastSeen    time.Time
}
