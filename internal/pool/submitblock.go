package pool

import (
	"log"
	"fmt"
	"github.com/techobg/prl-forge/internal/block"
	"encoding/hex"
	

)


func SubmitBlock(job *Job) error {


		

	log.Printf(
		"🚀 SubmitBlock job=%s height=%d proof=%d hs=%d",
		job.ID,
		job.Height,
		len(job.Proof),
		job.HS,
	)

	log.Printf("HeaderBytes : %d", len(job.HeaderBytes))
	log.Printf("MiningConfig: %d", len(job.MiningConfig))

	if job.ZKProof == nil {
		return fmt.Errorf("missing zk proof")
	}

	log.Printf("ZKProof OK")
	pb := &block.PearlBlock{
	Header: job.HeaderObj,
	Certificate:  job.Certificate,
	
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
log.Printf("Serialized block size: %d", len(raw))
if err := rpcClient.SubmitBlock(raw); err != nil {
	panic(err)
}

panic("RPC OK")

if err := rpcClient.SubmitBlock(raw); err != nil {
	log.Printf("❌ submitblock failed: %v", err)
	return err
}

log.Println("✅ submitblock accepted by RPC")





log.Printf("Serialized block: %d bytes", len(raw))

	

	return nil
}