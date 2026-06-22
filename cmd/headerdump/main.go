package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/techobg/prl-forge/internal/block"
	"github.com/techobg/prl-forge/internal/pearl"
)

func main() {
	client := pearl.New(pearl.RPCConfig{
		Host:     "127.0.0.1",
		Port:     44144, // или RPC портът на Pearl
		User:     "rpcuser",
		Password: "rpcpassword",
	})

	tpl, err := client.GetBlockTemplate()
	if err != nil {
		log.Fatal(err)
	}

	prevHash, err := hex.DecodeString(tpl.PreviousBlockHash)
	if err != nil {
		log.Fatal(err)
	}

	nbitsBytes, err := hex.DecodeString(tpl.Bits)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbitsBytes) != 4 {
		log.Fatalf("invalid nBits length: %d", len(nbitsBytes))
	}

	nbits := uint32(nbitsBytes[0])<<24 |
		uint32(nbitsBytes[1])<<16 |
		uint32(nbitsBytes[2])<<8 |
		uint32(nbitsBytes[3])

	merkleRoot, err := block.BuildMerkleRootFromTemplate(tpl)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	headerHex, err := header.Hex()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===== GBT =====")
fmt.Printf("Height   : %d\n", tpl.Height)
fmt.Printf("Version  : %08x\n", tpl.Version)
fmt.Printf("Bits     : %s\n", tpl.Bits)
fmt.Printf("PrevHash : %s\n", tpl.PreviousBlockHash)

fmt.Println()

fmt.Println("===== HEADER =====")
fmt.Printf("Header   : %s\n", headerHex)
}
