package pool

type CoinbaseHasher interface {
	Hash([]byte) ([]byte, error)
}

func BuildCoinbaseHash(
	hasher CoinbaseHasher,
	coinbase []byte,
) ([]byte, error) {
	return hasher.Hash(coinbase)
}