package stratum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"time"
  "strings"
  "strconv"

	"github.com/techobg/prl-forge/internal/block"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum/protocol"
	"github.com/techobg/prl-forge/internal/zkpow"
)

type SubmitParams struct {
	JobID      string  `json:"job_id"`
	PlainProof string  `json:"plain_proof"`
	HS         float64 `json:"hs"`
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

	jobID := params.JobID

if i := strings.IndexByte(jobID, '_'); i >= 0 {
	jobID = jobID[:i]
}

job, ok := GetJob(jobID)
	if !ok {
		return
	}

	job.Proof = append([]byte(nil), decoded...)
	job.HS = uint64(params.HS)
	job.Wallet = session.Wallet
	job.Worker = session.Worker

	if p := pool.Current(); p != nil {
		id := session.Wallet + "." + session.Worker

		if w := p.Workers().Get(id); w != nil {
			w.Shares++
			w.LastSeen = time.Now()
			w.Hashrate = params.HS
			p.History().Add(p.TotalHashrate())

			pool.WorkerHistory().Add(
				session.Wallet,
				session.Worker,
				int64(params.HS),
			)

			pool.MinerHistory().Add(
				session.Wallet,
				int64(params.HS),
			)

			log.Printf("WorkerHistory: %s.%s = %.0f", session.Wallet, session.Worker, params.HS)
		}

		// Track current mining round
		p.Round().AddShare(session.Difficulty)

		log.Printf(
			"ADD SHARE -> shares=%d work=%.2f diff=%.2f",
			p.Round().Shares(),
			p.Round().Work(),
			session.Difficulty,
		)

		// Persist current round
		if err := pool.SaveRoundStats(p.RoundHeight(), p.Round()); err != nil {
			log.Printf("failed to save round: %v", err)
		}
	}

	// Accept share immediately
	_ = session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	})

log.Println(">>> BEFORE ExtractZKProof")

	zk, err := zkpow.ExtractZKProof(job.HeaderBytes, job.Proof)
  
  log.Println(">>> AFTER ExtractZKProof")
  
	if err != nil {
		log.Printf("❌ ZK extract failed: %v", err)

		_ = session.Send(protocol.Response{
			ID:     req.ID,
			Result: false,
			Error:  []any{20, "Invalid ZK proof", nil},
		})
		return
	}

	job.ZKProof = zk
  
  
   
  bits, err := strconv.ParseUint(job.NBits, 16, 32)
if err != nil {
	log.Printf("parse nbits failed: %v", err)
} else {
	shareBits := pool.DifficultyToBitsFromNetwork(
	uint32(bits),
	session.Difficulty,
)

log.Printf(
	"VERIFY SHARE: networkBits=%08x shareBits=%08x diff=%.2f",
	uint32(bits),
	shareBits,
	session.Difficulty,
)

err = zkpow.VerifyWithNBits(
	job.HeaderBytes,
	job.ZKProof,
	shareBits,
)

	log.Printf("VerifyWithNBits: %v", err)
}
  
	cert, err := block.NewZKCertificate(
		job.HeaderObj,
		job.ZKProof,
		uint32(job.CertVersion),
	)
	if err != nil {
		log.Printf("❌ Certificate build failed: %v", err)
		return
	}

	job.Certificate = cert
  
  
  

	if err := zkpow.VerifyNetwork(job.HeaderBytes, job.ZKProof); err != nil {
	log.Printf("✅ Share accepted (not a block): %v", err)
	return
}

log.Println("🏆 BLOCK FOUND - submitting")

if err := pool.SubmitBlock(job); err != nil {
	log.Printf("❌ SubmitBlock failed: %v", err)
} else {
	log.Println("✅ SubmitBlock finished")
}

}