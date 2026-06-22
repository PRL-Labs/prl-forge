package pool

type Job struct {
	ID string

	// Block information
	Height int64
	Header string
	Target string

	PrevHash string
	Coinb1   string
	Coinb2   string
	Merkle   []string
	Version  string
	NBits    string
	NTime    string

	// Pearl certificate
	CertVersion int

	// Данни от BlockTemplate
	CoinbaseValue uint64
	CoinbaseFlags string

	// За бъдеща submit проверка
	Clean bool
}