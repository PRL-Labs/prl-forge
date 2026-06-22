package mining

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/techobg/prl-forge/internal/block"
	"github.com/techobg/prl-forge/internal/pool"
)

func TemplateFromJob(job *pool.Job) (*block.Template, error) {
	if job == nil {
		return nil, fmt.Errorf("nil job")
	}

	prevHash, err := hex.DecodeString(job.PrevHash)
	if err != nil {
		return nil, err
	}

	version, err := strconv.ParseUint(job.Version, 16, 32)
	if err != nil {
		return nil, err
	}

	bits, err := strconv.ParseUint(job.NBits, 16, 32)
	if err != nil {
		return nil, err
	}

	ntime, err := strconv.ParseUint(job.NTime, 16, 32)
	if err != nil {
		return nil, err
	}

	return &block.Template{
		Height:       uint64(job.Height),
		Version:      uint32(version),
		PreviousHash: prevHash,
		Bits:         uint32(bits),
		Timestamp:    uint32(ntime),
	}, nil
}
