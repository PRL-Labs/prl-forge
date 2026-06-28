package block

import (
	"bytes"
	"encoding/binary"
)

type TxOut struct {
	Value        uint64
	ScriptPubKey []byte
}

func (out *TxOut) Serialize(buf *bytes.Buffer) error {

	if err := binary.Write(buf, binary.LittleEndian, out.Value); err != nil {
		return err
	}

	if err := WriteVarInt(buf, uint64(len(out.ScriptPubKey))); err != nil {
		return err
	}

	buf.Write(out.ScriptPubKey)

	return nil
}