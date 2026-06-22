package block

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	IncompleteHeaderSize = 76
)

type IncompleteHeader struct {
	Version int32

	PrevBlock  []byte
	MerkleRoot []byte

	Timestamp uint32
	NBits     uint32
}

func (h *IncompleteHeader) Serialize() ([]byte, error) {
	if len(h.PrevBlock) != 32 {
		return nil, fmt.Errorf("prev block must be 32 bytes")
	}

	if len(h.MerkleRoot) != 32 {
		return nil, fmt.Errorf("merkle root must be 32 bytes")
	}

	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, h.Version); err != nil {
		return nil, err
	}

	buf.Write(h.PrevBlock)
	buf.Write(h.MerkleRoot)

	if err := binary.Write(&buf, binary.LittleEndian, h.Timestamp); err != nil {
		return nil, err
	}

	if err := binary.Write(&buf, binary.LittleEndian, h.NBits); err != nil {
		return nil, err
	}

	if buf.Len() != IncompleteHeaderSize {
		return nil, fmt.Errorf(
			"incomplete header size = %d expected %d",
			buf.Len(),
			IncompleteHeaderSize,
		)
	}

	return buf.Bytes(), nil
}
