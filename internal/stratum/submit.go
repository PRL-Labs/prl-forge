package stratum

import (
	"encoding/json"
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

type SubmitParams []string

func HandleSubmit(session *Session, req *protocol.Request) {
	log.Println("📤 mining.submit")

	var params SubmitParams

	if err := json.Unmarshal(req.Params, &params); err != nil {
		log.Printf("Invalid submit params: %v", err)
		return
	}

	if len(params) < 5 {
		log.Printf("Invalid submit request (%d params)", len(params))
		return
	}

	worker := params[0]
	jobID := params[1]
	extraNonce2 := params[2]
	nTime := params[3]
	nonce := params[4]

	log.Printf("Worker      : %s", worker)
	log.Printf("Job ID      : %s", jobID)
	log.Printf("ExtraNonce2 : %s", extraNonce2)
	log.Printf("NTime       : %s", nTime)
	log.Printf("Nonce       : %s", nonce)

	resp := protocol.Response{
		ID:     req.ID,
		Result: true,
		Error:  nil,
	}

	if err := session.Send(resp); err != nil {
		log.Println(err)
		return
	}

	log.Println("✅ Share accepted")
}