package tx

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

func (tx *Transaction) Serialize(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, tx.Version); err != nil {
		return err
	}

	if err := WriteVarInt(w, uint64(len(tx.Inputs))); err != nil {
		return err
	}

	for _, in := range tx.Inputs {
		if err := in.Serialize(w); err != nil {
			return err
		}
	}

	if err := WriteVarInt(w, uint64(len(tx.Outputs))); err != nil {
		return err
	}

	for _, out := range tx.Outputs {
		if err := out.Serialize(w); err != nil {
			return err
		}
	}

	return binary.Write(w, binary.LittleEndian, tx.LockTime)
}

func (tx *Transaction) Bytes() ([]byte, error) {
	var buf bytes.Buffer

	buf.Grow(tx.SerializeSize())

	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (tx *Transaction) TxID() ([32]byte, error) {
	var zero [32]byte

	raw, err := tx.Bytes()
	if err != nil {
		return zero, err
	}

	first := sha256.Sum256(raw)
	second := sha256.Sum256(first[:])

	return second, nil
}