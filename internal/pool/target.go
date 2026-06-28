package pool

import (
	"fmt"
	"math/big"
)

var powLimit = new(big.Int).Sub(
	new(big.Int).Lsh(big.NewInt(1), 208),
	big.NewInt(1),
)

func ShareTarget(diff uint64) string {
	if diff == 0 {
		diff = 1
	}

	target := new(big.Int).Div(
		new(big.Int).Set(powLimit),
		big.NewInt(int64(diff)),
	)

	return fmt.Sprintf("%064x", target)
}