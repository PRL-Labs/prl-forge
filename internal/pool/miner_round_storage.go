

 package pool

import (
    "encoding/json"
    "log"
    "os"
    
)


const minerRoundFile = "data/miner_round.json"

func SaveMinerRounds(m *MinerRoundManager) error {
	log.Println("=== SaveMinerRounds CALLED ===")

	_ = os.MkdirAll("data", 0755)

	f, err := os.Create(minerRoundFile)
	if err != nil {
		log.Printf("CREATE ERROR: %v", err)
		return err
	}
	defer f.Close()

	log.Printf("Writing miner rounds to %s", minerRoundFile)

	if err := json.NewEncoder(f).Encode(m.Export()); err != nil {
		log.Printf("ENCODE ERROR: %v", err)
		return err
	}

	log.Println("=== SaveMinerRounds DONE ===")
	return nil
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
        log.Printf(
                "RESTORE wallet=%s shares=%d work=%.2f",
                wallet,
                state.Shares,
                state.Work,
        )

       m.rounds[wallet] = &RoundStats{
    shares: state.Shares,
    work:   state.Work,
}
}

	return m, nil
}