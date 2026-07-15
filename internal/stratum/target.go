package stratum

import (
	"fmt"
	"math/big"
)

var diff1Target = func() *big.Int {
	t := new(big.Int)
	t.SetString("00000000ffff0000000000000000000000000000000000000000000000000000", 16)
	return t
}()

func DifficultyToTarget(diff float64) string {
	if diff <= 0 {
		diff = 1
	}

	d := big.NewInt(int64(diff))

	target := new(big.Int).Div(new(big.Int).Set(diff1Target), d)

	return fmt.Sprintf("%064x", target)
}