package block

import "fmt"

func Build(t *Template) (*Header, error) {
	if t == nil {
		return nil, fmt.Errorf("nil template")
	}

	ih := &IncompleteHeader{
		Version: int32(t.Version),

		PrevBlock: append([]byte(nil), t.PreviousHash...),

		MerkleRoot: append([]byte(nil), t.MerkleRoot...),

		Timestamp: t.Timestamp,

		NBits: t.Bits,
	}

	raw, err := ih.Serialize()
	if err != nil {
		return nil, err
	}

	return &Header{
		IncompleteHeader: raw,
		ProofCommitment:  make([]byte, ProofCommitmentSize),
	}, nil
}