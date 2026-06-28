package block

import (
	"bytes"
	"encoding/binary"
)

type ZKCertificate struct {
	CertVersion uint32
	HeaderHash  []byte
	PublicData  []byte
	ProofData   []byte
}

func (z *ZKCertificate) ProofCommitment() []byte {
	buf := make([]byte, 4+len(z.PublicData))

	binary.LittleEndian.PutUint32(buf[:4], z.CertVersion)
	copy(buf[4:], z.PublicData)

	return DoubleSHA256(buf)
}

func (z *ZKCertificate) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	if err := binary.Write(&buf, binary.LittleEndian, z.CertVersion); err != nil {
		return nil, err
	}

	buf.Write(z.HeaderHash)

	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(z.PublicData))); err != nil {
		return nil, err
	}

	buf.Write(z.PublicData)

	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(z.ProofData))); err != nil {
		return nil, err
	}

	buf.Write(z.ProofData)

	return buf.Bytes(), nil
}