package block

import (
	"encoding/hex"
	"fmt"
)

const (
	ProofCommitmentSize = 32
)

type Header struct {
	// Точно както PearlHeader.incomplete_header
	IncompleteHeader []byte

	// 32-byte commitment
	ProofCommitment []byte
}

func (h *Header) SerializeWithoutProofCommitment() []byte {
	out := make([]byte, len(h.IncompleteHeader))
	copy(out, h.IncompleteHeader)
	return out
}

func (h *Header) Serialize() ([]byte, error) {
	if len(h.ProofCommitment) != ProofCommitmentSize {
		return nil, fmt.Errorf(
			"invalid proof commitment size: expected %d, got %d",
			ProofCommitmentSize,
			len(h.ProofCommitment),
		)
	}

	out := make([]byte, 0, len(h.IncompleteHeader)+ProofCommitmentSize)
	out = append(out, h.IncompleteHeader...)
	out = append(out, h.ProofCommitment...)

	return out, nil
}

func (h *Header) Hex() (string, error) {
	data, err := h.Serialize()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}
