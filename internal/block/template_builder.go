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
fmt.Println("============== TX DEBUG ==============")

for i, tx := range tpl.Transactions {
	if i >= 3 {
		break
	}

	fmt.Printf("TX[%d]\n", i)
	fmt.Printf("  TXID : %s\n", tx.TxID)
	fmt.Printf("  HASH : %s\n", tx.Hash)
}

fmt.Println("======================================")
	fmt.Println("============== MERKLE DEBUG ==============")
	fmt.Printf("CoinbaseHash : %x\n", cbHash)
	fmt.Printf("Height       : %d\n", tpl.Height)
	fmt.Printf("CoinbaseValue: %d\n", tpl.CoinbaseValue)
	fmt.Printf("Flags        : %s\n", tpl.CoinbaseAux.Flags)
	fmt.Printf("TX count     : %d\n", len(tpl.Transactions))

	hashes := make([][]byte, 0, len(tpl.Transactions)+1)
	hashes = append(hashes, cbHash)

	for _, tx := range tpl.Transactions {
		h, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return nil, fmt.Errorf("invalid tx hash: %w", err)
		}

		hashes = append(hashes, h)
	}

	fmt.Printf("Merkle leaves: %d\n", len(hashes))
	fmt.Println("=========================================")

	return BuildMerkleRoot(hashes), nil
}