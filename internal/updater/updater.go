package updater

import (
	"log"
	"time"
"crypto/sha256"
	"github.com/techobg/prl-forge/internal/pearl"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum"
	"github.com/techobg/prl-forge/internal/stats"
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
	log.Println("UPDATER START")
	go func() {
		for {

			template, err := u.rpc.GetBlockTemplate()
if err != nil {
				log.Printf("Updater: getblocktemplate failed: %v", err)
				time.Sleep(u.interval)
				continue
			}
			log.Printf("RequiredCertVersion = %d", template.RequiredCertVersion)

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

			job := u.engine.BuildJob(template)

			stats.Update(template)

sum := sha256.Sum256(job.HeaderBytes)

log.Println("******** NEW JOB ********")
log.Printf("JobID        : %s", job.ID)
log.Printf("Height       : %d", job.Height)
log.Printf("Header SHA   : %x", sum)


			log.Printf("ENGINE (Updater): %p", u.engine)

			// Engine вече пази CurrentJob.
			// Няма повече stratum.SetCurrentJob()

			if u.stratum != nil {
				u.stratum.Broadcast(job)
			}

			time.Sleep(u.interval)
		}
	}()
}
