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
	password := params[1]

	session.Authorized = true
	session.Wallet = wallet
	session.Worker = wallet // временно
	session.Difficulty = 1

	log.Printf("Wallet: %s", wallet)
	log.Printf("Password: %s", password)

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
		[]any{1.0},
	); err != nil {
		log.Println("Notify error:", err)
	} else {
		log.Println("Difficulty OK")
	}

	job := NewDummyJob()

	log.Println("Sending job...")

	if err := SendJob(session, job); err != nil {
		log.Println("Job error:", err)
	} else {
		log.Println("Job OK")
	}
}