package block

import (
	"bytes"
	"encoding/binary"
)

func WriteVarInt(buf *bytes.Buffer, v uint64) error {

	switch {

	case v < 0xfd:
		return buf.WriteByte(byte(v))

	case v <= 0xffff:

		buf.WriteByte(0xfd)
		return binary.Write(buf, binary.LittleEndian, uint16(v))

	case v <= 0xffffffff:

		buf.WriteByte(0xfe)
		return binary.Write(buf, binary.LittleEndian, uint32(v))

	default:

		buf.WriteByte(0xff)
		return binary.Write(buf, binary.LittleEndian, v)
	}
}