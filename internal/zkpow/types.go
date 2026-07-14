package zkpow

type ZKProof struct {
	PublicData []byte
	ProofData  []byte
        HashJackpot []byte
}

var LastProof *ZKProof
