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

	log.Printf("Worker      : %s", worker)
	log.Printf("Job ID      : %s", jobID)
	log.Printf("ExtraNonce2 : %s", extraNonce2)
	log.Printf("NTime       : %s", nTime)
	log.Printf("Nonce       : %s", nonce)

	job, ok := jobManager.Get(jobID)
	if !ok {
		log.Printf("❌ Unknown job: %s", jobID)

		resp := protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{21, "Job not found", nil},
		}

		_ = session.Send(resp)
		return
	}

	log.Printf("✅ Job %s found", job.ID)

	share := &Share{
		Worker:      worker,
		Wallet:      session.Wallet,
		JobID:       jobID,
		ExtraNonce2: extraNonce2,
		NTime:       nTime,
		Nonce:       nonce,
		Difficulty:  session.Difficulty,
		Accepted:    true,
		Time:        time.Now(),
	}

	shareManager.Add(share)

	log.Printf("📊 Total shares: %d", shareManager.Count())

	resp := protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	}

	if err := session.Send(resp); err != nil {
		log.Println(err)
		return
	}

	log.Println("✅ Share accepted")
}