
	package pool

import "log"

var roundLoaded bool

func (p *Pool) RestoreRound(height int64) {

	if roundLoaded {
		return
	}

	r, savedHeight, err := LoadRoundStats()
if err == nil {
	p.round = r
	p.roundHeight = savedHeight
  
  if minerRounds, err := LoadMinerRounds(); err == nil {
	p.minerRounds = minerRounds
}

	log.Printf(
		"✅ Restored round: height=%d shares=%d",
		savedHeight,
		r.Shares(),
	)
}

	roundLoaded = true
}

func (p *Pool) StartNewRound(height int64) {

	p.round.Reset()
	p.roundHeight = height
  
  
  p.minerRounds = NewMinerRoundManager()

if err := SaveMinerRounds(p.minerRounds); err != nil {
	log.Printf("failed to save miner rounds: %v", err)
}

	if err := SaveRoundStats(height, p.round); err != nil {
		log.Printf("failed to save round: %v", err)
	}

	log.Printf("🏁 New pool round: height=%d", height)
}