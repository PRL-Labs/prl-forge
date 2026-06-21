package pool

import (
	"fmt"

	"github.com/techobg/prl-forge/internal/pearl"
)

type Builder struct {
	nextID uint64
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(tpl *pearl.BlockTemplate) *Job {
	b.nextID++

	job := &Job{
		ID:             fmt.Sprintf("%d", b.nextID),
		Height:         tpl.Height,
		PrevHash:       tpl.PreviousBlockHash,
		Version:        fmt.Sprintf("%08x", uint32(tpl.Version)),
		NBits:          tpl.Bits,
		NTime:          fmt.Sprintf("%08x", uint32(tpl.CurTime)),
		Target:         tpl.Target,
		CoinbaseValue:  tpl.CoinbaseValue,
		CoinbaseFlags:  tpl.CoinbaseAux.Flags,
		ExtraNonce1:    fmt.Sprintf("%08x", b.nextID),
		ExtraNonce2Size: 4,
		Coinb1:         "01000000",
		Coinb2:         "ffffffff",
		Merkle:         make([]string, 0, len(tpl.Transactions)),
		Clean:          true,
	}

	for _, tx := range tpl.Transactions {
		job.Merkle = append(job.Merkle, tx.Hash)
	}

	return job
}