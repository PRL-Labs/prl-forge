package pool

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/techobg/prl-forge/internal/block"
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

	// Decode previous block hash (HEX -> []byte)
	prevHash, err := hex.DecodeString(tpl.PreviousBlockHash)
	if err != nil {
		panic(fmt.Errorf("invalid previous block hash: %w", err))
	}

	// Decode nBits (HEX -> uint32)
	bitsBytes, err := hex.DecodeString(tpl.Bits)
	if err != nil {
		panic(fmt.Errorf("invalid bits: %w", err))
	}

	if len(bitsBytes) != 4 {
		panic(fmt.Errorf("invalid bits length: %d", len(bitsBytes)))
	}

	nbits := binary.BigEndian.Uint32(bitsBytes)

	merkleRoot, err := block.BuildMerkleRootFromTemplate(tpl)
	if err != nil {
		panic(err)
	}

	bt := &block.Template{
		Height:       uint64(tpl.Height),
		Version:      uint32(tpl.Version),
		PreviousHash: prevHash,
		MerkleRoot:   merkleRoot,
		Timestamp:    uint32(tpl.CurTime),
		Bits:         nbits,
	}

	header, err := block.Build(bt)
	if err != nil {
		panic(err)
	}

	headerHex, err := header.Hex()
	if err != nil {
		panic(err)
	}
fmt.Printf("Header bytes: %d\n", len(headerHex)/2)
fmt.Println("RPC Target :", tpl.Target)
	job := &Job{
	
		ID:            fmt.Sprintf("%d", b.nextID),
		Height:        tpl.Height,
		Header:        headerHex,
		Target:        tpl.Target,
		CertVersion:   2,
		CoinbaseValue: tpl.CoinbaseValue,
		CoinbaseFlags: tpl.CoinbaseAux.Flags,
		Clean:         true,
	}

	return job
}
