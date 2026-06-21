package stratum

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func ValidateShare(share *Share, job *Job) error {

	// 1. basic validation
	if err := validateExtraNonce2(share.ExtraNonce2); err != nil {
		return err
	}
	if err := validateNTime(share.NTime); err != nil {
		return err
	}
	if err := validateNonce(share.Nonce); err != nil {
		return err
	}

	// 2. build header
	header, err := buildBitcoinHeader(share, job)
	if err != nil {
		return fmt.Errorf("header build failed: %w", err)
	}

	// 3. sha256d
	hash := sha256d(header)
	hash = reverseBytes(hash)

	// 4. convert to bigint
	hashInt := new(big.Int).SetBytes(hash)

	// 5. FIXED: SetString returns 2 values
	target := new(big.Int)
	_, ok := target.SetString(job.Target, 16)
	if !ok {
		return fmt.Errorf("invalid target format")
	}

	if target.Sign() <= 0 {
		return fmt.Errorf("invalid target value")
	}

	// 6. compare
	if hashInt.Cmp(target) > 0 {
		return fmt.Errorf("share does not meet target")
	}

	return nil
}

// --------------------
// HEADER BUILD
// --------------------

func buildBitcoinHeader(share *Share, job *Job) ([]byte, error) {

	version := hexToLE(job.Version)
	prevHash := hexToLE(job.PrevHash)

	merkleRoot := hexToLE(job.Merkle[0])

	ntime := hexToLE(share.NTime)
	bits := hexToLE(job.NBits)
	nonce := hexToLE(share.Nonce)

	header := make([]byte, 0, 80)

	header = append(header, version...)    // 4 bytes
	header = append(header, prevHash...)   // 32 bytes
	header = append(header, merkleRoot...) // 32 bytes
	header = append(header, ntime...)      // 4 bytes
	header = append(header, bits...)       // 4 bytes
	header = append(header, nonce...)      // 4 bytes

	return header, nil
}

// --------------------
// CRYPTO
// --------------------

func sha256d(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

func reverseBytes(b []byte) []byte {
	out := make([]byte, len(b))
	for i := range b {
		out[i] = b[len(b)-1-i]
	}
	return out
}

// --------------------
// HEX helper
// --------------------

func hexToLE(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		return []byte{}
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	return b
}

// --------------------
// BASIC VALIDATION
// --------------------

func validateExtraNonce2(v string) error {
	if len(v) != 8 {
		return fmt.Errorf("invalid extranonce2 length")
	}
	_, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("invalid extranonce2")
	}
	return nil
}

func validateNTime(v string) error {
	if len(v) != 8 {
		return fmt.Errorf("invalid ntime length")
	}
	_, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("invalid ntime")
	}
	return nil
}

func validateNonce(v string) error {
	if len(v) != 8 {
		return fmt.Errorf("invalid nonce length")
	}
	_, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("invalid nonce")
	}
	return nil
}