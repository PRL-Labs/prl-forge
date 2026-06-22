package block

import (
	"bytes"
	"encoding/binary"
)

type Coinbase struct {
	Height uint64
	Value  uint64
	Flags  string
}

func BuildCoinbase(c Coinbase) ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, c.Height); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, c.Value); err != nil {
		return nil, err
	}

	buf.WriteString(c.Flags)

	return buf.Bytes(), nil
}

func CoinbaseHash(c Coinbase) ([]byte, error) {
	cb, err := BuildCoinbase(c)
	if err != nil {
		return nil, err
	}

	return DoubleSHA256(cb), nil
}