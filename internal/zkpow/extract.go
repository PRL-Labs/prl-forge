package zkpow

import "fmt"

func ExtractZKProof(header []byte, proof []byte) (*ZKProof, error) {
	if err := ProvePlain(header, proof); err != nil {
		return nil, err
	}

	if LastProof == nil {
		return nil, fmt.Errorf("zk proof is nil")
	}

	return LastProof, nil
}