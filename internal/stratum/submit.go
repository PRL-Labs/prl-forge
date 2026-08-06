package stratum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"time"
 /// "strings"
  "strconv"
  "math/big"

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

///if i := strings.IndexByte(jobID, '_'); i >= 0 {
///	jobID = jobID[:i]
///}

job, ok := GetJob(jobID)
if !ok {
        log.Printf("JOB NOT FOUND: %s", jobID)
        return
}

log.Printf("JOB PTR=%p ID=%s HEIGHT=%d", job, job.ID, job.Height)

	job.Proof = append([]byte(nil), decoded...)
	job.HS = uint64(params.HS)
	job.Wallet = session.Wallet
	job.Worker = session.Worker

	 p := pool.Current()
if p == nil {
    return
}
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


	

	// Accept share immediately

	_ = session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	})

///p.Round().AddWork(shareWork)
///p.MinerRounds().AddWork(session.Wallet, shareWork)

newDiff := pool.VarDiff().ObserveShare(session.Wallet)

session.Difficulty = newDiff

if err := SendCurrentJob(session); err != nil {
	log.Printf("SendCurrentJob: %v", err)
}

if err := pool.SaveRoundStats(p.RoundHeight(), p.Round()); err != nil {
	log.Printf("failed to save round: %v", err)
}

if err := pool.SaveMinerRounds(p.MinerRounds()); err != nil {
	log.Printf("failed to save miner rounds: %v", err)
}


log.Println(">>> BEFORE ExtractZKProof")

	log.Println("1")

zk, err := zkpow.ExtractZKProof(job.HeaderBytes, job.Proof)

log.Println("2")

log.Printf(">>> AFTER ExtractZKProof err=%v", err)
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
  
  
 
   
bits64, err := strconv.ParseUint(job.NBits, 16, 32)
if err != nil {
	log.Printf("parse nbits failed: %v", err)
} else {
	networkBits := uint32(bits64)

	

shareBits := pool.DifficultyToBitsFromNetwork(
        networkBits,
        job.Difficulty,
)


log.Printf(
    "VERIFY JOBDIFF=%.0f SESSIONDIFF=%.0f",
    job.Difficulty,
    session.Difficulty,
)

testBits := pool.DifficultyToBitsFromNetwork(
    networkBits,
    8,
)

log.Printf(
    "TEST DIFF=8 -> %.8f",
    pool.BitsToDifficulty(testBits),
)

	log.Printf(
		"VERIFY SHARE: networkBits=%08x shareBits=%08x diff=%.2f",
		networkBits,
		shareBits,
		session.Difficulty,
	)
  
  shareDifficulty := pool.BitsToDifficulty(shareBits)

log.Printf(
    "SHARE DIFFICULTY = %.8f",
    shareDifficulty,
)
  
  shareWork := pool.CalcWork(shareBits)
networkWork := pool.CalcWork(networkBits)

ratio := new(big.Rat).SetFrac(shareWork, networkWork)
log.Printf("WORK RATIO = %s", ratio.FloatString(12))


log.Printf(
    "SHARE WORK=%s NETWORK WORK=%s",
    shareWork.String(),
    networkWork.String(),
)
  




log.Println(">>> BEFORE VerifyWithNBits")



log.Printf("JOB PTR VERIFY=%p ID=%s HEIGHT=%d", job, job.ID, job.Height)

log.Println("A")

err = zkpow.VerifyWithNBits(
job.HeaderBytes,
job.ZKProof,
shareBits,
)

log.Println("B")

log.Printf(">>> AFTER VerifyWithNBits err=%v", err)

err2 := zkpow.VerifyNetwork(job.HeaderBytes, job.ZKProof)

log.Printf("COMPARE -> share=%v network=%v", err, err2)

if err != nil {
	log.Printf("VerifyWithNBits failed: %v", err)
	return
}



log.Printf("VerifyWithNBits OK")



  
	cert, err := block.NewZKCertificate(
		job.HeaderObj,
		job.ZKProof,
		uint32(job.CertVersion),
	)
	if err != nil {
		log.Printf("❌ Certificate build failed: %v", err)
		
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
}