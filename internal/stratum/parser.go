package stratum

import (
	"encoding/json"
	"log"

	"github.com/techobg/prl-forge/internal/stratum/protocol"
)

func (s *Session) dispatch(line string) {
	var req protocol.Request

	if err := json.Unmarshal([]byte(line), &req); err != nil {
		log.Printf("Invalid JSON: %v", err)
		return
	}

	Dispatch(s, &req)
}