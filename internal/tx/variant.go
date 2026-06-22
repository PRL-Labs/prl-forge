package tx

import (
	"encoding/binary"
	"io"
)

func WriteVarInt(w io.Writer, val uint64) error {
	switch {
	case val < 0xfd:
		_, err := w.Write([]byte{byte(val)})
		return err

	case val <= 0xffff:
		if _, err := w.Write([]byte{0xfd}); err != nil {
			return err
		}
		return binary.Write(w, binary.LittleEndian, uint16(val))

	case val <= 0xffffffff:
		if _, err := w.Write([]byte{0xfe}); err != nil {
			return err
		}
		return binary.Write(w, binary.LittleEndian, uint32(val))

	default:
		if _, err := w.Write([]byte{0xff}); err != nil {
			return err
		}
		return binary.Write(w, binary.LittleEndian, val)
	}
}

func VarIntSerializeSize(val uint64) int {
	switch {
	case val < 0xfd:
		return 1
	case val <= 0xffff:
		return 3
	case val <= 0xffffffff:
		return 5
	default:
		return 9
	}
}