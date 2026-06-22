package pool

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/techobg/prl-forge/internal/block"
)

func (j *Job) ToTemplate() (*block.Template, error) {
	if j == nil {
		return nil, fmt.Errorf("nil job")
	}

	prevHash, err := hex.DecodeString(j.PrevHash)
	if err != nil {
		return nil, fmt.Errorf("decode prevhash: %w", err)
	}

	version, err := strconv.ParseUint(j.Version, 16, 32)
	if err != nil {
		return nil, fmt.Errorf("parse version: %w", err)
	}

	bits, err := strconv.ParseUint(j.NBits, 16, 32)
	if err != nil {
		return nil, fmt.Errorf("parse bits: %w", err)
	}

	ntime, err := strconv.ParseUint(j.NTime, 16, 32)
	if err != nil {
		return nil, fmt.Errorf("parse ntime: %w", err)
	}

	return &block.Template{
		Height:       uint64(j.Height),
		Version:      uint32(version),
		PreviousHash: prevHash,
		Timestamp:    uint32(ntime),
		Bits:         uint32(bits),

		// TODO: ще се попълни от Merkle Builder
		MerkleRoot: make([]byte, 32),
	}, nil
}