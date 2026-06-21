package updater

import (
	"log"
	"time"

	"github.com/techobg/prl-forge/internal/pearl"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum"
)

type Updater struct {
	rpc      *pearl.Client
	engine   *pool.Engine
	stratum  *stratum.Server
	interval time.Duration

	lastPrevHash string
	lastTarget   string
	lastHeight   int64
}

func New(
	rpcClient *pearl.Client,
	engine *pool.Engine,
	stratumServer *stratum.Server,
	interval time.Duration,
) *Updater {

	if interval <= 0 {
		interval = 10 * time.Second
	}

	return &Updater{
		rpc:      rpcClient,
		engine:   engine,
		stratum:  stratumServer,
		interval: interval,
	}
}

func (u *Updater) Start() {
	go func() {

		for {

			template, err := u.rpc.GetBlockTemplate()
			if err != nil {
				log.Printf("Updater: getblocktemplate failed: %v", err)
				time.Sleep(u.interval)
				continue
			}

			// skip if same chain state
			if template.PreviousBlockHash == u.lastPrevHash &&
				template.Target == u.lastTarget &&
				template.Height == u.lastHeight {

				time.Sleep(u.interval)
				continue
			}

			u.lastPrevHash = template.PreviousBlockHash
			u.lastTarget = template.Target
			u.lastHeight = template.Height

			log.Println("========================================")
			log.Printf("Height       : %d", template.Height)
			log.Printf("Version      : %d", template.Version)
			log.Printf("PrevHash     : %s", template.PreviousBlockHash)
			log.Printf("Bits         : %s", template.Bits)
			log.Printf("Target       : %s", template.Target)
			log.Printf("Coinbase     : %d", template.CoinbaseValue)
			log.Printf("Flags        : %s", template.CoinbaseAux.Flags)
			log.Printf("Transactions : %d", len(template.Transactions))
			log.Println("========================================")

			// build job
			job := u.engine.BuildJob(template)

			// IMPORTANT: register job globally for submit lookup
			stratum.SetCurrentJob(job)

			// broadcast to miners
			if u.stratum != nil {
				u.stratum.Broadcast(job)
			}

			log.Printf(
				"📦 Broadcasted job id=%s height=%d prev=%s",
				job.ID,
				template.Height,
				template.PreviousBlockHash,
			)

			time.Sleep(u.interval)
		}
	}()
}