package stratum

import (
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

func HandleSubscribe(session *Session, req *protocol.Request) {
	log.Println("⛏️ mining.subscribe")

	session.Subscribed = true

	// REAL Stratum response format
	extraNonce1 := "00000001"
	extraNonce2Size := 4

	resp := protocol.Response{
		ID: req.ID,
		Result: []interface{}{
			[]interface{}{
				[]interface{}{"mining.set_difficulty", "1"},
				[]interface{}{"mining.notify", "1"},
			},
			extraNonce1,
			extraNonce2Size,
		},
		Error: nil,
	}

	if err := session.Send(resp); err != nil {
		log.Println(err)
		return
	}

	// Send initial job immediately
	if err := SendCurrentJob(session); err != nil {
		log.Println(err)
	}
}