package block

import (
	"crypto/sha256"
)

// DoubleSHA256 изпълнява SHA256(SHA256(data)).
func DoubleSHA256(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])

	hash := make([]byte, len(second))
	copy(hash, second[:])

	return hash
}
