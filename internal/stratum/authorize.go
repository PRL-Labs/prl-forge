package stratum

import (
	"encoding/json"
	"log"
	"strings"

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

	log.Printf("Miner authorized")
	log.Printf("Wallet : %s", wallet)
	log.Printf("Worker : %s", worker)

	// authorize response
	if err := session.Send(protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	}); err != nil {
		log.Printf("authorize response: %v", err)
		return
	}

	// difficulty
	if err := session.Notify(
		"mining.set_difficulty",
		[]any{session.Difficulty},
	); err != nil {
		log.Printf("set_difficulty: %v", err)
		return
	}

	// first job
	if err := SendCurrentJob(session); err != nil {
		log.Printf("notify: %v", err)
		return
	}

	log.Println("✅ Authorization completed")
}
