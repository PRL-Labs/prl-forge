package pool

import (
	"encoding/json"
	"os"
  "math/big"
)

const roundFile = "data/round.json"

type roundState struct {
    Height int64  `json:"height"`
    Shares uint64 `json:"shares"`
    Work   string `json:"work"`
}

func SaveRoundStats(height int64, r *RoundStats) error {
	r.mu.RLock()
	state := roundState{
    Height: height,
    Shares: r.shares,
    Work:   r.work.String(),
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

work := new(big.Int)

if state.Work != "" {
    if _, ok := work.SetString(state.Work, 10); !ok {
        work = new(big.Int)
    }
}

r := &RoundStats{
    shares:    state.Shares,
    work:      work,
    lastWork:  new(big.Int),
}

	return r, state.Height, nil
}
