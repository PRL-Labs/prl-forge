package stratum

import (
	"encoding/json"
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type AuthorizeParams []string

func HandleAuthorize(session *Session, req *protocol.Request) {
	log.Println("🔐 mining.authorize")

	var params AuthorizeParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		log.Printf("Invalid authorize params: %v", err)
		return
	}

	if len(params) < 2 {
		log.Println("Invalid authorize request")
		return
	}

	wallet := params[0]
	_ = params[1]

	// worker parsing (wallet.worker)
	worker := wallet
	if idx := len(wallet) - 1; idx > 0 {
		worker = wallet
	}

	session.Authorized = true
	session.Wallet = wallet
	session.Worker = worker
	session.Difficulty = 1.0

	log.Printf("Wallet: %s", wallet)

	resp := protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	}

	if err := session.Send(resp); err != nil {
		log.Println(err)
		return
	}

	log.Println("Sending difficulty...")

	if err := session.Notify(
		"mining.set_difficulty",
		[]any{float64(session.Difficulty)},
	); err != nil {
		log.Println(err)
		return
	}

	log.Println("Difficulty OK")

	log.Println("Sending job...")

	if err := SendCurrentJob(session); err != nil {
		log.Println(err)
		return
	}

	log.Println("Job OK")
}