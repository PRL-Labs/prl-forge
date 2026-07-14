package stratum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"time"

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

	job, ok := GetJob(params.JobID)
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
pool.ActivityHistory().Add(
        session.Wallet,
        session.Difficulty,
)

newDiff := pool.VarDiff().ObserveShare(session.Wallet)

log.Printf(
        "VARDIFF wallet=%s current=%.0f suggested=%.0f",
        session.Wallet,
        session.Difficulty,
        newDiff,
)

if newDiff != session.Difficulty {

        log.Printf(
                "VARDIFF UPDATE %.0f -> %.0f",
                session.Difficulty,
                newDiff,
        )

        session.Difficulty = newDiff

        if err := SendDifficulty(session, newDiff); err != nil {
                log.Printf("SendDifficulty: %v", err)
        }

        if err := SendCurrentJob(session); err != nil {
                log.Printf("SendCurrentJob: %v", err)
        }
}





			log.Printf("WorkerHistory: %s.%s = %.0f", session.Wallet, session.Worker, params.HS)
		}

		// Track current mining round

p.Round().AddShare(session.Difficulty)

p.MinerRounds().AddShare(session.Wallet, session.Difficulty)

log.Printf(
    "ROUND HEIGHT SAVE -> %d",
    p.RoundHeight(),
)

if err := pool.SaveRoundStats(p.RoundHeight(), p.Round()); err != nil {
    log.Printf("failed to save round: %v", err)
}

if err := pool.SaveMinerRounds(p.MinerRounds()); err != nil {
	log.Printf("failed to save miner rounds: %v", err)
}


      


		log.Printf(
			"ADD SHARE -> shares=%d work=%.2f diff=%.2f",
			p.Round().Shares(),
			p.Round().Work(),
			session.Difficulty,
		)

		// Persist current round
    
    log.Printf(
    "ROUND HEIGHT SAVE -> %d",
    p.RoundHeight(),
)
    
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

	zk, err := zkpow.ExtractZKProof(job.HeaderBytes, job.Proof)
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

log.Printf("HEADER COMMITMENT = %x", job.HeaderObj.ProofCommitment)
log.Printf("CERT COMMITMENT   = %x", cert.ProofCommitment())
log.Printf("HEADER HASH       = %x", job.Certificate.HeaderHash)
  
  log.Printf("HEADER COMMITMENT = %x", job.HeaderObj.ProofCommitment)
log.Printf("CERT COMMITMENT   = %x", cert.ProofCommitment())
  

if err := pool.SubmitBlock(job); err != nil {
	log.Printf("❌ SubmitBlock failed: %v", err)
} else {
	log.Println("✅ SubmitBlock finished")

	if p := pool.Current(); p != nil {
	p.StartNewRound(job.Height + 1)
	p.MinerRounds().Reset(job.Wallet)

	if err := pool.SaveMinerRounds(p.MinerRounds()); err != nil {
		log.Printf("failed to save miner rounds: %v", err)
	}
}
 }
 
}