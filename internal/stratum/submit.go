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
		log.Printf("submit: invalid params: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid submit parameters", nil},
		})
		return
	}

	if len(params) < 5 {
		log.Printf("submit: expected 5 params, got %d", len(params))

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid submit parameters", nil},
		})
		return
	}

	worker := params[0]
	jobID := params[1]
	extraNonce2 := params[2]
	nTime := params[3]
	nonce := params[4]

	job, ok := GetJob(jobID)
	if !ok {
		log.Printf("submit: unknown job %s", jobID)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{21, "Job not found", nil},
		})
		return
	}

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

	if err := ValidateShare(share, job); err != nil {
		log.Printf("submit: share rejected: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid share", nil},
		})
		return
	}

	share.Accepted = true
	shareManager.Add(share)

	log.Printf(
		"✅ Share accepted worker=%s job=%s total=%d",
		worker,
		jobID,
		shareManager.Count(),
	)

	_ = session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	})
}
