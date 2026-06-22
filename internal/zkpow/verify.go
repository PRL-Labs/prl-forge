//go:build zkpow

package zkpow

import "fmt"

func Verify(header []byte, proof []byte) error {
	rc := verify(header, proof)

	switch rc {
	case 0:
		return nil

	case 1:
		return fmt.Errorf("invalid proof")

	default:
		return fmt.Errorf("zk verifier returned %d", rc)
	}
}