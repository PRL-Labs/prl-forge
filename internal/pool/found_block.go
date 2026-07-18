package pool

import (
	"log"
	"time"

	blocks "github.com/techobg/prl-forge/internal/updater/blocks"
)

func ProcessFoundBlock(job *Job, wallet string) error {
	log.Printf("?? ProcessFoundBlock called (wallet=%s height=%d)", wallet, job.Height)

	if p := Current(); p != nil {

		p.Blocks().Add(&blocks.Block{
			Height:    job.Height,
			Hash:      "",
			Finder:    wallet,
			Reward:    0,
			Timestamp: time.Now(),
		})

		log.Printf("?? Resetting pool round")
		p.Round().Reset()

		log.Printf("?? Resetting miner round: %s", wallet)
		p.MinerRounds().Reset(wallet)
	}

	return nil
}