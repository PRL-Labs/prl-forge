package pool

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/techobg/prl-forge/internal/block"
)

func SubmitBlock(job *Job) error {

	log.Printf(
		"🚀 SubmitBlock job=%s height=%d",
		job.ID,
		job.Height,
	)

	if job.ZKProof == nil {
		return fmt.Errorf("missing zk proof")
	}

	pb := &block.PearlBlock{
		Header:      job.HeaderObj,
		Certificate: job.Certificate,
	}

	if job.Template != nil {
		for _, tx := range job.Template.Transactions {
			rawTx, err := hex.DecodeString(tx.Data)
			if err != nil {
				return err
			}

			pb.Transactions = append(pb.Transactions, rawTx)
		}
	}

	raw, err := pb.Serialize()
	if err != nil {
		return err
	}

	if err := rpcClient.SubmitBlock(raw); err != nil {
		log.Printf("❌ submitblock failed: %v", err)
		return err
	}

	log.Printf("🏆 Block accepted by Pearl RPC (height=%d)", job.Height)

	return nil
}
