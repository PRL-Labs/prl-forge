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
fmt.Printf("PrevBlock : %x\n", h.PrevBlock)
fmt.Printf("MerkleRoot: %x\n", h.MerkleRoot)
	prev := make([]byte, 32)
copy(prev, h.PrevBlock)

merkle := make([]byte, 32)
copy(merkle, h.MerkleRoot)

// reverse byte order
for i := 0; i < 16; i++ {
	prev[i], prev[31-i] = prev[31-i], prev[i]
	merkle[i], merkle[31-i] = merkle[31-i], merkle[i]
}

buf.Write(prev)
buf.Write(merkle)

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
