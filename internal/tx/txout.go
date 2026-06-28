package tx

import (
	"encoding/binary"
	"io"
)

type TxOut struct {
	Value    int64
	PkScript Script
}

func NewTxOut(value int64, pkScript Script) *TxOut {
	return &TxOut{
		Value:    value,
		PkScript: pkScript,
	}
}

func (t *TxOut) Serialize(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, t.Value); err != nil {
		return err
	}

	return t.PkScript.Serialize(w)
}

func (t *TxOut) SerializeSize() int {
	return 8 + t.PkScript.Size()
}
