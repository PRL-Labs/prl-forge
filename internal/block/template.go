package block

type Template struct {
	Height uint64

	Version uint32

	PreviousHash []byte
	MerkleRoot   []byte

	Timestamp uint32
	Bits      uint32

	ExtraNonce []byte

	CoinbaseValue uint64
	CoinbaseFlags string
}