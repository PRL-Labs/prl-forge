package tx

import (
	"encoding/binary"
	"io"
)

const (
	MaxPrevOutIndex = 0xffffffff
	MaxSequence     = 0xffffffff
)

type OutPoint struct {
	Hash  [32]byte
	Index uint32
}

type TxIn struct {
    PreviousOutPoint OutPoint
    SignatureScript  Script
    Sequence         uint32

    // SegWit witness stack
    Witness []Script
}


func NewCoinbaseInput(script Script) *TxIn {
	return &TxIn{
		PreviousOutPoint: OutPoint{
			Hash:  [32]byte{},
			Index: MaxPrevOutIndex,
		},
		SignatureScript: script,
		Sequence:        MaxSequence,
	}
}

func (o OutPoint) Serialize(w io.Writer) error {
	if _, err := w.Write(o.Hash[:]); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, o.Index)
}

func (t *TxIn) Serialize(w io.Writer) error {
	if err := t.PreviousOutPoint.Serialize(w); err != nil {
		return err
	}

	if err := t.SignatureScript.Serialize(w); err != nil {
		return err
	}

	return binary.Write(w, binary.LittleEndian, t.Sequence)
}

func (t *TxIn) SerializeSize() int {
	return 32 +
		4 +
		t.SignatureScript.Size() +
		4
}
