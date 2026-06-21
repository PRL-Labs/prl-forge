package pool

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func BuildCoinbase(
	coinb1 string,
	extraNonce1 string,
	extraNonce2 string,
	coinb2 string,
) ([]byte, error) {

	var buf bytes.Buffer

	b1, err := decodeHex(coinb1)
	if err != nil {
		return nil, err
	}
	buf.Write(b1)

	en1, err := decodeHex(extraNonce1)
	if err != nil {
		return nil, err
	}
	buf.Write(en1)

	en2, err := decodeHex(extraNonce2)
	if err != nil {
		return nil, err
	}
	buf.Write(en2)

	b2, err := decodeHex(coinb2)
	if err != nil {
		return nil, err
	}
	buf.Write(b2)

	return buf.Bytes(), nil
}

func decodeHex(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, fmt.Errorf("invalid hex length")
	}

	dst := make([]byte, hex.DecodedLen(len(s)))
	_, err := hex.Decode(dst, []byte(s))
	if err != nil {
		return nil, fmt.Errorf("invalid coinbase hex: %w", err)
	}

	return dst, nil
}