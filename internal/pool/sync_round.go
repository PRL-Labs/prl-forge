package pool

import "log"

var roundLoaded bool

func (p *Pool) SyncRound(height int64) {

	if !roundLoaded {

		r, savedHeight, err := LoadRoundStats()
		if err == nil {

			if savedHeight == height {
				p.round = r
				p.roundHeight = height
				log.Printf("✅ Restored round: height=%d shares=%d", height, r.Shares())
			}

		}

		roundLoaded = true
	}

	if p.roundHeight != height {

		p.round.Reset()
		p.roundHeight = height

		if err := SaveRoundStats(height, p.round); err != nil {
			log.Printf("failed to save round: %v", err)
		}

		log.Printf("🔄 New mining round: height=%d", height)
	}
}
