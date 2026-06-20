package stratum

import (
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

func HandleSubscribe(session *Session, req *protocol.Request) {
	log.Println("⛏️ mining.subscribe")
    session.Subscribed = true
	resp := protocol.Response{
		ID: req.ID,
		Result: []interface{}{
			[]interface{}{
				[]interface{}{"mining.set_difficulty", "1"},
				[]interface{}{"mining.notify", "1"},
			},
			"PRLForge-Session",
			4,
		},
		Error: nil,
	}

	if err := session.Send(resp); err != nil {
		log.Println(err)
	}
}