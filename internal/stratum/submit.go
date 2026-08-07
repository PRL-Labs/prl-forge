package stratum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
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
			ID: req.ID, Result: false,
			Error: []any{20, "Invalid submit parameters", nil},
		})
		return
	}

	// Decode proof
	proof, err := base64.StdEncoding.DecodeString(params.PlainProof)
	if err != nil {
		log.Printf("proof decode error: %v", err)
		return
	}

	zr, err := gzip.NewReader(bytes.NewReader(proof))
	if err != nil {
		log.Printf("gzip error: %v", err)
		return
	}
	defer zr.Close()

	decoded, err := io.ReadAll(zr)
	if err != nil {
		log.Printf("gzip read error: %v", err)
		return
	}

	// Clean jobID
	jobID := params.JobID
	if i := strings.IndexByte(jobID, '_'); i >= 0 {
		jobID = jobID[:i]
	}

	job, ok := GetJob(jobID)
	if !ok {
		log.Printf("JOB NOT FOUND: %s", jobID)
		return
	}

	job.Proof = append([]byte(nil), decoded...)
	job.HS = uint64(params.HS)
	job.Wallet = session.Wallet
	job.Worker = session.Worker

	p := pool.Current()
	if p == nil {
		return
	}

	// Worker update
	id := session.Wallet + "." + session.Worker
	if w := p.Workers().Get(id); w != nil {
		w.Shares++
		w.LastSeen = time.Now()
		w.Hashrate = params.HS
    
    now := time.Now()

if !w.LastShareTime.IsZero() {
	delta := now.Sub(w.LastShareTime).Seconds()

	targetTime := 22.0 // 🔥 тук си прав — 22-25 е по-добре

	if delta < targetTime/2 {
		w.Difficulty *= 1.5
	} else if delta > targetTime*2 {
		w.Difficulty *= 0.7
	}

	if w.Difficulty < 1 {
		w.Difficulty = 1
	}
}

w.LastShareTime = now

		p.History().Add(p.TotalHashrate())

		pool.WorkerHistory().Add(session.Wallet, session.Worker, int64(params.HS))
		pool.MinerHistory().Add(session.Wallet, int64(params.HS))
	}

// =========================
// ✅ ACCEPT SHARE
// =========================
_ = session.Send(protocol.Response{
	ID: req.ID, Result: true, Error: nil,
})

// =========================
// ✅ WORK CALC (CORRECT)
// =========================

// parse nBits (hex → uint32)
bits64, err := strconv.ParseUint(job.NBits, 16, 32)
if err != nil {
	log.Printf("parse nbits failed: %v", err)
	return
}

networkBits := uint32(bits64)

// =========================
// 🔥 1. Bits → Target
// =========================
networkTarget := pool.BitsToTarget(networkBits)

// =========================
// 🔥 2. Target / Difficulty
// =========================
targetFloat := new(big.Float).SetInt(networkTarget)

// ⚠️ CRITICAL: difficulty must be float64 > 0
if session.Difficulty <= 0 {
	session.Difficulty = 1
}

targetFloat.Quo(targetFloat, big.NewFloat(session.Difficulty))

shareTarget := new(big.Int)
targetFloat.Int(shareTarget)

// =========================
// 🔥 3. Target → Work
// =========================
bits64, err = strconv.ParseUint(job.NBits, 16, 32)
if err != nil {
	log.Printf("parse nbits failed: %v", err)
	return
}

networkBits = uint32(bits64)


shareWork := new(big.Int).SetUint64(uint64(session.Difficulty * 4294967296))

log.Printf(
	"SHARE DEBUG -> diff=%.2f networkBits=%08x shareWork=%s",
	session.Difficulty,
	networkBits,
	shareWork.String(),
)

// =========================
// ✅ ADD WORK → LUCK
// =========================
p.Round().AddWork(shareWork)
p.MinerRounds().AddWork(session.Wallet, shareWork)
p.Round().AddDifficulty(session.Difficulty)

// =========================
// 💾 SAVE
// =========================
_ = pool.SaveRoundStats(p.RoundHeight(), p.Round())
_ = pool.SaveMinerRounds(p.MinerRounds())
	// =========================
	// ZK VERIFY
	// =========================
	zk, err := zkpow.ExtractZKProof(job.HeaderBytes, job.Proof)
	if err != nil {
		log.Printf("ZK extract failed: %v", err)
		return
	}
	job.ZKProof = zk

	err = zkpow.VerifyWithNBits(
		job.HeaderBytes,
		job.ZKProof,
		networkBits,
	)
	if err != nil {
		log.Printf("VerifyWithNBits failed: %v", err)
		return
	}

	// =========================
	// BLOCK CHECK
	// =========================
	cert, err := block.NewZKCertificate(
		job.HeaderObj,
		job.ZKProof,
		uint32(job.CertVersion),
	)
	if err != nil {
		log.Printf("Certificate error: %v", err)
	}

	job.Certificate = cert

	if err := zkpow.VerifyNetwork(job.HeaderBytes, job.ZKProof); err != nil {
		log.Println("Share OK (not block)")
		return
	}

	log.Println("🏆 BLOCK FOUND")

	if err := pool.SubmitBlock(job); err != nil {
		log.Printf("SubmitBlock failed: %v", err)
	}
}