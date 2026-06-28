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
fmt.Printf("TX count: %d\n", len(tpl.Transactions))
if len(tpl.Transactions) > 0 {
    fmt.Printf("First TX hash : %s\n", tpl.Transactions[0].Hash)
fmt.Printf("First TX txid : %s\n", tpl.Transactions[0].TxID)
}
fmt.Printf("CoinbaseValue: %d\n", tpl.CoinbaseValue)
fmt.Printf("CoinbaseFlags: %s\n", tpl.CoinbaseAux.Flags)
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
		
        CoinbaseValue: tpl.CoinbaseValue,
        CoinbaseFlags: tpl.CoinbaseAux.Flags,
	}


	header, err := block.Build(bt)
	if err != nil {
		panic(err)
	}

	headerHex, err := header.Hex()
	if err != nil {
		panic(err)
	}

	headerBytes := header.SerializeWithoutProofCommitment()
	fmt.Println("========== HEADER DEBUG ==========")


fmt.Printf("Header len : %d\n", len(headerBytes))
fmt.Printf("Header hex : %x\n", headerBytes)

fmt.Println("RPC Target :", tpl.Target)
fmt.Println("RPC Bits   :", tpl.Bits)
fmt.Println("==============================")
	fmt.Printf("Header len: %d\n", len(headerBytes))
fmt.Printf("Header hex: %x\n", headerBytes)
fmt.Printf("JobID: %s\n", fmt.Sprintf("%d", b.nextID))
fmt.Printf("Header bytes: %d\n", len(headerHex)/2)
fmt.Printf("HeaderBytes len: %d\n", len(headerBytes))
fmt.Println("RPC Target :", tpl.Target)
fmt.Println("RPC Bits   :", tpl.Bits)
fmt.Printf("headerHex          : %s\n", headerHex)
fmt.Printf("headerBytes->hex   : %x\n", headerBytes)
target := ShareTarget(1)
	
job := &Job{
	ID:            fmt.Sprintf("%d", b.nextID),
	Height:        tpl.Height,
	Header:        headerHex,
	HeaderObj:     header,
    HeaderBytes:   headerBytes,
	Target:        target,
	CertVersion:   2,
	CoinbaseValue: tpl.CoinbaseValue,
	CoinbaseFlags: tpl.CoinbaseAux.Flags,
	Template:      tpl,
	Clean:         true,
}


	return job

}