package tx

import (
	"io"
)

type Script []byte

func NewScript(data []byte) Script {
	if len(data) == 0 {
		return nil
	}

	s := make([]byte, len(data))
	copy(s, data)
	return Script(s)
}

func (s Script) Bytes() []byte {
	if s == nil {
		return nil
	}

	b := make([]byte, len(s))
	copy(b, s)
	return b
}

func (s Script) Size() int {
	return VarIntSerializeSize(uint64(len(s))) + len(s)
}

func (s Script) Serialize(w io.Writer) error {
	if err := WriteVarInt(w, uint64(len(s))); err != nil {
		return err
	}

	if len(s) == 0 {
		return nil
	}

	_, err := w.Write(s)
	return err
}
