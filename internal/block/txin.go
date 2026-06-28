package block

import (
	"bytes"
	"encoding/binary"
)

type TxIn struct {
	PrevHash  [32]byte
	PrevIndex uint32
	ScriptSig []byte
	Sequence  uint32

	// SegWit
	Witness [][]byte
}

func (in *TxIn) Serialize(buf *bytes.Buffer) error {

	buf.Write(in.PrevHash[:])

	if err := binary.Write(buf, binary.LittleEndian, in.PrevIndex); err != nil {
		return err
	}

	if err := WriteVarInt(buf, uint64(len(in.ScriptSig))); err != nil {
		return err
	}

	buf.Write(in.ScriptSig)

	return binary.Write(buf, binary.LittleEndian, in.Sequence)
}