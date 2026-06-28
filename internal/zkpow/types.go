package zkpow

type ZKProof struct {
	PublicData []byte
	ProofData  []byte
}

var LastProof *ZKProof