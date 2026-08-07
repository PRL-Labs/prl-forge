package pool

import (
	"math/big"
)

// Bitcoin-style max target (difficulty 1)
var maxTarget = func() *big.Int {
	target := new(big.Int)
	target.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
	return target
}()

// Bits -> Target
func BitsToTarget(bits uint32) *big.Int {
	exponent := bits >> 24
	mantissa := bits & 0xFFFFFF

	target := new(big.Int).SetUint64(uint64(mantissa))

	shift := 8 * (exponent - 3)
	target.Lsh(target, uint(shift))

	return target
}

// Target -> Difficulty
func BitsToDifficulty(bits uint32) float64 {
	target := BitsToTarget(bits)

	rat := new(big.Rat).SetFrac(maxTarget, target)
	f, _ := rat.Float64()

	return f
}

// Difficulty -> Target
func DifficultyToTarget(diff float64) *big.Int {
	rat := new(big.Rat).SetFloat64(diff)

	targetRat := new(big.Rat).Quo(new(big.Rat).SetInt(maxTarget), rat)

	target := new(big.Int)
	targetRat.Num().Div(targetRat.Num(), targetRat.Denom())

	target.Set(targetRat.Num())

	return target
}

// Network bits + share diff -> share bits
func DifficultyToBitsFromNetwork(networkBits uint32, shareDiff float64) uint32 {
	networkDiff := BitsToDifficulty(networkBits)

	realDiff := networkDiff * shareDiff

	target := DifficultyToTarget(realDiff)

	return TargetToBits(target)
}

// Target -> Bits
func TargetToBits(target *big.Int) uint32 {
	bytes := target.Bytes()

	size := len(bytes)

	var compact uint32

	if size <= 3 {
		for i := 0; i < size; i++ {
			compact <<= 8
			compact |= uint32(bytes[i])
		}
		compact <<= 8 * uint32(3-size)
	} else {
		compact = uint32(bytes[0])<<16 | uint32(bytes[1])<<8 | uint32(bytes[2])
	}

	if compact&0x00800000 != 0 {
		compact >>= 8
		size++
	}

	compact |= uint32(size) << 24

	return compact
}

// Work = 2^256 / (target+1)
func CalcWork(bits uint32) *big.Int {
	target := BitsToTarget(bits)

	one := big.NewInt(1)
	target.Add(target, one)

	two256 := new(big.Int).Lsh(one, 256)

	return new(big.Int).Div(two256, target)
}
func TargetToWork(target *big.Int) *big.Int {
	if target.Sign() <= 0 {
		return big.NewInt(0)
	}

	// work = 2^256 / target
	two256 := new(big.Int).Lsh(big.NewInt(1), 256)

	return new(big.Int).Div(two256, target)
}