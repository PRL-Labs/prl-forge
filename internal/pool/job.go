package pool

type Job struct {
	ID string

	Height int64

	PrevHash string

	Coinb1 string
	Coinb2 string

	Merkle []string

	Version string
	NBits   string
	NTime   string

	// Данни от Pearl BlockTemplate
	Target         string
	CoinbaseValue  uint64
	CoinbaseFlags  string

	// Подготовка за Stratum
	ExtraNonce1     string
	ExtraNonce2Size int

	Clean bool
}