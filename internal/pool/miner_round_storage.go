package pool

import (
	"encoding/json"
	"os"
)

const minerRoundFile = "data/miner_round.json"

func SaveMinerRounds(m *MinerRoundManager) error {
	_ = os.MkdirAll("data", 0755)

	f, err := os.Create(minerRoundFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(m.Export())
}

func LoadMinerRounds() (*MinerRoundManager, error) {
	m := NewMinerRoundManager()

	f, err := os.Open(minerRoundFile)
	if err != nil {
		return m, err
	}
	defer f.Close()

	var data map[string]MinerRoundData
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return m, err
	}

	for wallet, state := range data {
		m.rounds[wallet] = &RoundStats{
			shares: state.Shares,
			work:   state.Work,
		}
	}

	return m, nil
}