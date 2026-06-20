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

	return &Job{
		ID:       fmt.Sprintf("%d", b.nextID),
		PrevHash: tpl.PreviousBlockHash,
		Version:  fmt.Sprintf("%08x", uint32(tpl.Version)),
		NBits:    tpl.Bits,
		NTime:    fmt.Sprintf("%08x", uint32(tpl.CurTime)),
		Coinb1:   "01000000",
		Coinb2:   "ffffffff",
		Merkle:   []string{},
		Clean:    true,
	}
}