package blocks

import "time"

type Block struct {
	Height    int64     `json:"height"`
	Hash      string    `json:"hash"`
	Finder    string    `json:"finder"`
	Reward    float64   `json:"reward"`
	Timestamp time.Time `json:"timestamp"`
}