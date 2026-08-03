 package pool

import "math/big"

var (
	bigOne    = big.NewInt(1)
	oneLsh256 = new(big.Int).Lsh(bigOne, 256)
)

func CompactToBig(compact uint32) *big.Int {
	mantissa := compact & 0x007fffff
	isNegative := compact&0x00800000 != 0
	exponent := uint(compact >> 24)

	var bn *big.Int

	if exponent <= 3 {
		mantissa >>= 8 * (3 - exponent)
		bn = big.NewInt(int64(mantissa))
	} else {
		bn = big.NewInt(int64(mantissa))
		bn.Lsh(bn, 8*(exponent-3))
	}

	if isNegative {
		bn.Neg(bn)
	}

	return bn
}

func BigToCompact(n *big.Int) uint32 {
	if n.Sign() == 0 {
		return 0
	}

	var mantissa uint32

	exponent := uint(len(n.Bytes()))

	if exponent <= 3 {
		mantissa = uint32(n.Bits()[0])
		mantissa <<= 8 * (3 - exponent)
	} else {
		tmp := new(big.Int).Set(n)
		mantissa = uint32(tmp.Rsh(tmp, 8*(exponent-3)).Bits()[0])
	}

	if mantissa&0x00800000 != 0 {
		mantissa >>= 8
		exponent++
	}

	compact := uint32(exponent<<24) | mantissa

	if n.Sign() < 0 {
		compact |= 0x00800000
	}

	return compact
}

func DifficultyToBitsFromNetwork(networkBits uint32, shareDiff float64) uint32 {
	if shareDiff <= 1 {
		return networkBits
	}

	networkTarget := CompactToBig(networkBits)

	targetFloat := new(big.Float).SetInt(networkTarget)
	targetFloat.Mul(targetFloat, big.NewFloat(shareDiff))

	target, _ := targetFloat.Int(nil)

	if target.Sign() <= 0 {
		target = big.NewInt(1)
	}

	return BigToCompact(target)
}


func BitsToDifficulty(bits uint32) float64 {
	powLimitBits := uint32(0x1b00ffff)

	max := CompactToBig(powLimitBits)
	target := CompactToBig(bits)

	difficulty := new(big.Rat).SetFrac(max, target)

	diff, _ := difficulty.Float64()

	return diff
}


// CalcWork  1:1  Pearl node.
func CalcWork(bits uint32) *big.Int {
	target := CompactToBig(bits)

	if target.Sign() <= 0 {
		return big.NewInt(0)
	}

	denominator := new(big.Int).Add(target, bigOne)

	return new(big.Int).Div(oneLsh256, denominator)
}