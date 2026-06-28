package block

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/techobg/prl-forge/internal/zkpow"
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

	switch z.CertVersion {
	case 1: // ZK_DENSE
		buf.Write(z.PublicData)

	case 2: // ZK_MOE
		if err := binary.Write(&buf, binary.LittleEndian, uint32(len(z.PublicData))); err != nil {
			return nil, err
		}
		buf.Write(z.PublicData)

	default:
		return nil, fmt.Errorf("unsupported certificate version %d", z.CertVersion)
	}

	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(z.ProofData))); err != nil {
		return nil, err
	}

	buf.Write(z.ProofData)

	return buf.Bytes(), nil
}

func NewZKCertificate(
	header *Header,
	proof *zkpow.ZKProof,
	certVersion uint32,
) (*ZKCertificate, error) {

	headerBytes, err := header.Serialize()
	if err != nil {
		return nil, err
	}


	

	
	return &ZKCertificate{
		CertVersion: certVersion,
		HeaderHash:  DoubleSHA256(headerBytes),
		PublicData:  proof.PublicData,
		ProofData:   proof.ProofData,
	}, nil
}