package pool

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func BuildMerkleRoot(coinbaseHash []byte, branches []string) ([]byte, error) {
	root := make([]byte, len(coinbaseHash))
	copy(root, coinbaseHash)

	for _, branch := range branches {
		hash, err := hex.DecodeString(branch)
		if err != nil {
			return nil, fmt.Errorf("invalid merkle branch: %w", err)
		}

		data := append(root, hash...)

		first := sha256.Sum256(data)
		second := sha256.Sum256(first[:])

		root = second[:]
	}

	return root, nil
}
