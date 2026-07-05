package stratum

import (
	"encoding/json"
	"log"
	"strings"
     "net"
	"time"

	"github.com/techobg/prl-forge/internal/pool"
	workers "github.com/techobg/prl-forge/internal/updater/workers"
	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type AuthorizeParams []string

type AuthorizeV2Params struct {
	Wallet string `json:"wallet"`
	Agent  string `json:"agent"`
	Type   string `json:"type"`
}

func HandleAuthorize(session *Session, req *protocol.Request) {
	log.Println("🔐 mining.authorize")

	var wallet string

	// Stratum V1
	var v1 AuthorizeParams
	if err := json.Unmarshal(req.Params, &v1); err == nil && len(v1) > 0 {
		wallet = v1[0]
	} else {
		// SRBMiner V2
		var v2 AuthorizeV2Params

		if err := json.Unmarshal(req.Params, &v2); err != nil {
			log.Printf("authorize: invalid params: %v", err)
			return
		}

		wallet = v2.Wallet
	}

	if wallet == "" {
		log.Println("authorize: empty wallet")
		return
	}

	var worker string

	parts := strings.SplitN(wallet, ".", 2)
	if len(parts) == 2 {
		wallet = parts[0]
		worker = parts[1]
	}

	session.Authorized = true
	session.Wallet = wallet
	session.Worker = worker
	session.Difficulty = 1.0

	if p := pool.Current(); p != nil {

	ip := ""

	if addr, ok := session.conn.RemoteAddr().(*net.TCPAddr); ok {
		ip = addr.IP.String()
	}

	p.Workers().Add(&workers.Worker{
		ID:          wallet + "." + worker,
		Wallet:      wallet,
		Name:        worker,
		IP:          ip,
		ConnectedAt: time.Now(),
		LastSeen:    time.Now(),
	})
}

	log.Println("Miner authorized")
	log.Printf("Wallet : %s", wallet)
	log.Printf("Worker : %s", worker)

	// Authorize response
	if err := session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
		Type:   "v2",
	}); err != nil {
		log.Printf("authorize response: %v", err)
		return
	}

	// Pearl V2 (SRBMiner/Kryptex) НЕ изпраща mining.set_difficulty
	// Веднага след authorize изпращаме първия job.

	if err := SendCurrentJob(session); err != nil {
		log.Printf("notify: %v", err)
		return
	}

	log.Println("✅ Authorization completed")
}