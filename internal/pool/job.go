package pool

import (
    
    "github.com/techobg/prl-forge/internal/zkpow"
	"github.com/techobg/prl-forge/internal/block"
	"github.com/techobg/prl-forge/internal/pearl"
)






 type Job struct {
	ID string

	// Block information
Wallet string
	Worker string

	Height int64
	Header string
	HeaderBytes  []byte
    MiningConfig []byte
    
    
	Target string
  Difficulty float64
  
    HeaderObj *block.Header
	Certificate *block.ZKCertificate
PearlBlock  *block.PearlBlock
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

	// Original GBT
    Template *pearl.BlockTemplate

	// Cached submit data
Proof []byte
HS uint64

ProofHash string
ZKProof *zkpow.ZKProof

	// Cached full block (RC2)
   

	// За бъдеща submit проверка
	Clean bool

    

}