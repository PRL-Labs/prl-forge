package stratum

import (
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

func HandleSubscribe(session *Session, req *protocol.Request) {
	log.Println("⛏️ mining.subscribe")

	session.Subscribed = true

	const extraNonce1 = "00000001"
	const extraNonce2Size = 4

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
		log.Printf("subscribe response error: %v", err)
		return
	}

	log.Println("✅ mining.subscribe completed")
}
