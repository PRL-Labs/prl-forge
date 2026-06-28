package block

import (
	"bytes"
)

type PearlBlock struct {
	Header       *Header
	Certificate  *ZKCertificate
	Transactions [][]byte
}

func (b *PearlBlock) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	certBytes, err := b.Certificate.Serialize()
	if err != nil {
		return nil, err
	}

	headerBytes, err := b.Header.Serialize()
	if err != nil {
		return nil, err
	}

	buf.Write(certBytes)
	buf.Write(headerBytes)

	if err := WriteVarInt(&buf, uint64(len(b.Transactions))); err != nil {
		return nil, err
	}

	for _, tx := range b.Transactions {
		buf.Write(tx)
	}

	return buf.Bytes(), nil
}
