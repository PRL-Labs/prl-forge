package updater

import (
	"log"
	"time"

	"github.com/techobg/prl-forge/internal/pearl"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stats"
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
	log.Println("Updater started")

	go func() {
		for {

			template, err := u.rpc.GetBlockTemplate()
			if err != nil {
				log.Printf("Updater: getblocktemplate failed: %v", err)
				time.Sleep(u.interval)
				continue
			}

			if template.PreviousBlockHash == u.lastPrevHash &&
				template.Target == u.lastTarget &&
				template.Height == u.lastHeight {

				time.Sleep(u.interval)
				continue
			}

			u.lastPrevHash = template.PreviousBlockHash
			u.lastTarget = template.Target
			u.lastHeight = template.Height

			job := u.engine.BuildJob(template)

			stats.Update(template)

			if p := pool.Current(); p != nil {
				p.SyncRound(template.Height)

				log.Printf("📈 Pool hashrate: %d", p.TotalHashrate())

			}

			if u.stratum != nil {
				u.stratum.Broadcast(job)
			}

			time.Sleep(u.interval)
		}
	}()
}
