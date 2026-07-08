package stratum

import "github.com/techobg/prl-forge/internal/stratum/protocol"

func SendDifficulty(session *Session, diff float64) error {
	return session.Send(protocol.Notification{
		ID:     nil,
		Method: "mining.set_difficulty",
		Params: []any{diff},
	})
}
