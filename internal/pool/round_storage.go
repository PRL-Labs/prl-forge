package pool

import (
	"encoding/json"
	"os"
)

const roundFile = "data/round.json"

type roundState struct {
	Height int64   `json:"height"`
	Shares uint64  `json:"shares"`
	Work   float64 `json:"work"`
}

func SaveRoundStats(height int64, r *RoundStats) error {
	r.mu.RLock()
	state := roundState{
		Height: height,
		Shares: r.shares,
		Work:   r.work,
	}
	r.mu.RUnlock()

	_ = os.MkdirAll("data", 0755)

	f, err := os.Create(roundFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(&state)
}

func LoadRoundStats() (*RoundStats, int64, error) {
	f, err := os.Open(roundFile)
	if err != nil {
		return NewRoundStats(), 0, err
	}
	defer f.Close()

	var state roundState
	if err := json.NewDecoder(f).Decode(&state); err != nil {
		return NewRoundStats(), 0, err
	}

	r := &RoundStats{
		shares: state.Shares,
		work:   state.Work,
	}

	return r, state.Height, nil
}
