package stratum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
    "github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type SubmitParams struct {
	JobID      string `json:"job_id"`
	PlainProof string `json:"plain_proof"`
	HS         uint64 `json:"hs"`
}

func HandleSubmit(session *Session, req *protocol.Request) {
	log.Println("📤 mining.submit")

	var params SubmitParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		log.Printf("submit decode error: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid submit parameters", nil},
		})
		return
	}

	log.Printf("JobID      : %s", params.JobID)
	log.Printf("HS         : %d", params.HS)
	log.Printf("Proof chars: %d", len(params.PlainProof))

	// Base64 decode
	proof, err := base64.StdEncoding.DecodeString(params.PlainProof)
	if err != nil {
		log.Printf("proof decode error: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid proof", nil},
		})
		return
	}

	log.Printf("Proof bytes: %d", len(proof))

	// Gzip decompress
	zr, err := gzip.NewReader(bytes.NewReader(proof))
	if err != nil {
		log.Printf("gzip error: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid gzip proof", nil},
		})
		return
	}
	defer zr.Close()

	decoded, err := io.ReadAll(zr)
	if err != nil {
		log.Printf("gzip read error: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid proof payload", nil},
		})
		return
	}

	log.Printf("Decoded proof: %d bytes", len(decoded))

	n := 64
	if len(decoded) < n {
		n = len(decoded)
	}

	log.Printf("Proof prefix: %x", decoded[:n])

	job, ok := GetJob(params.JobID)
if ok {
	job.Proof = append([]byte(nil), decoded...)
	job.HS = params.HS
	log.Printf("📦 Cached proof: %d bytes HS=%d", len(job.Proof), job.HS)
}

	// DEBUG ONLY
	_ = session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	})

	log.Println("✅ DEBUG submit accepted")
	if ok {
	_ = pool.SubmitBlock(job)
}
}