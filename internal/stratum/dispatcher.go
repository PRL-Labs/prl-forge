package stratum

import (
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

func Dispatch(session *Session, req *protocol.Request) {
	log.Printf("📩 RPC: %s", req.Method)

	switch req.Method {

	case "mining.subscribe":
		HandleSubscribe(session, req)

	case "mining.authorize":
		HandleAuthorize(session, req)

	case "mining.submit":
		HandleSubmit(session, req)

	default:
		log.Printf("⚠️ Unknown method: %s", req.Method)
	}
}
