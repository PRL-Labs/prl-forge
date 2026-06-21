package stratum

import (
	"encoding/json"
	"log"
	"time"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type SubmitParams []string

func HandleSubmit(session *Session, req *protocol.Request) {
	log.Println("📤 mining.submit")

	var params SubmitParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		log.Printf("Invalid submit params: %v", err)
		return
	}

	if len(params) < 5 {
		log.Printf("Invalid submit request (%d params)", len(params))
		return
	}

	worker := params[0]
	jobID := params[1]
	extraNonce2 := params[2]
	nTime := params[3]
	nonce := params[4]

	job, ok := GetJob(jobID)
	if !ok {
		log.Printf("❌ Unknown job: %s", jobID)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{21, "Job not found", nil},
		})
		return
	}

	log.Printf("Worker: %s Job: %s", worker, jobID)

	share := &Share{
		Worker:      worker,
		Wallet:      session.Wallet,
		JobID:       jobID,
		ExtraNonce2: extraNonce2,
		NTime:       nTime,
		Nonce:       nonce,
		Difficulty:  session.Difficulty,
		Time:        time.Now(),
	}

	// REAL validation gate
	if err := ValidateShare(share, job); err != nil {
		log.Printf("❌ Share rejected: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid share", nil},
		})
		return
	}

	share.Accepted = true
	shareManager.Add(share)

	log.Printf("📊 Shares: %d", shareManager.Count())

	_ = session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	})

	log.Println("✅ Share accepted")
}