package blockwatcher

import (
	"log"
	"time"

	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/updater/blocks"
)

type Watcher struct {
	lastHeight  int64
	poolAddress string
}

func New(poolAddress string) *Watcher {
	return &Watcher{
		poolAddress: poolAddress,
	}
}

func (w *Watcher) Start() {
	go w.loop()
}

func (w *Watcher) loop() {

	log.Println("🧱 BlockWatcher started")
	log.Println("Pool address:", w.poolAddress)

	for {

		client := pool.Client()
		if client == nil {
			time.Sleep(5 * time.Second)
			continue
		}

		height, err := client.GetBlockCount()
		if err != nil {
			log.Println("BlockWatcher:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if height != w.lastHeight {

			log.Printf("🧱 New chain height: %d", height)

			hash, err := client.GetBlockHash(height)
			if err != nil {
				log.Println("GetBlockHash:", err)
			} else {

				log.Println("Block hash:", hash)

				block, err := client.GetBlock(hash)
				if err != nil {
					log.Println("GetBlock:", err)
				} else {

					if len(block.RawTx) > 0 &&
						len(block.RawTx[0].Vout) > 0 {

						coinbase := block.RawTx[0].Vout[0].ScriptPubKey.Address

						log.Println("Coinbase address:", coinbase)

						if coinbase == w.poolAddress {

							log.Println("🎉 PRL Forge found a block!")

							if p := pool.Current(); p != nil {

								p.Blocks().Add(&blocks.Block{
									Height:    block.Height,
									Hash:      block.Hash,
									Finder:    coinbase,
									Reward:    0,
									Timestamp: time.Unix(block.Time, 0),
								})
							}
						}
					}
				}
			}

			w.lastHeight = height
		}

		time.Sleep(5 * time.Second)
	}
}
