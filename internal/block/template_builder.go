package block

import (
	"encoding/hex"
	"fmt"

	"github.com/techobg/prl-forge/internal/pearl"
)

func BuildMerkleRootFromTemplate(tpl *pearl.BlockTemplate) ([]byte, error) {
	cbHash, err := CoinbaseHash(Coinbase{
		Height: uint64(tpl.Height),
		Value:  tpl.CoinbaseValue,
		Flags:  tpl.CoinbaseAux.Flags,
	})
	if err != nil {
		return nil, err
	}

	hashes := make([][]byte, 0, len(tpl.Transactions)+1)

	hashes = append(hashes, cbHash)

	for _, tx := range tpl.Transactions {
		h, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return nil, fmt.Errorf("invalid tx hash: %w", err)
		}

		hashes = append(hashes, h)
	}

	return BuildMerkleRoot(hashes), nil
}