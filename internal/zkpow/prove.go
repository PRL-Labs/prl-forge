//go:build zkpow

package zkpow

/*
#cgo windows CFLAGS: -ID:/Projects/prl-forge/pearl-master/zk-pow/bindings/go
#cgo windows LDFLAGS: -LD:/Projects/prl-forge/pearl-master/zk-pow/bindings/go/target/x86_64-pc-windows-gnu/release -l:libzk_pow_ffi.dll.a

#cgo linux CFLAGS: -I/opt/prl-forge/pearl-master/zk-pow/bindings/go
#cgo linux LDFLAGS: -L/opt/prl-forge/pearl-master/zk-pow/bindings/go/target/release -lzk_pow_ffi

#include <stdlib.h>
#include <string.h>
#include "zk_pow_ffi.h"
*/
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"
  "log"
)

const miningConfigSize = 52

var defaultMiningConfigV1 = [miningConfigSize]byte{
	0x00, 0x04, 0x00, 0x00,
	0x20, 0x00,
	0x00, 0x00,
	0x07, 0x01, 0x03, 0x01, 0x00, 0x00,
	0x00, 0x01, 0x03, 0x01, 0x01, 0x01,
}

func headerToC(header []byte) C.IncompleteBlockHeader {

log.Printf("HEADER LEN=%d", len(header))
log.Printf("HEADER HEX=%x", header)

	var cHeader C.IncompleteBlockHeader

	cHeader.version = C.uint32_t(binary.LittleEndian.Uint32(header[0:4]))

	for i := 0; i < 32; i++ {
		cHeader.prev_block[i] = C.uint8_t(header[35-i])
		cHeader.merkle_root[i] = C.uint8_t(header[67-i])
	}

	cHeader.timestamp = C.uint32_t(binary.LittleEndian.Uint32(header[68:72]))
	cHeader.nbits = C.uint32_t(binary.LittleEndian.Uint32(header[72:76]))
  
  
  log.Printf(
    "C HEADER version=%08x time=%08x nbits=%08x",
    uint32(cHeader.version),
    uint32(cHeader.timestamp),
    uint32(cHeader.nbits),
)

	return cHeader
}

func ProvePlain(header []byte, proof []byte) error {
	if len(header) == 0 || len(proof) == 0 {
		return fmt.Errorf("invalid input")
	}

	miningConfig := defaultMiningConfigV1[:]

	_ = unsafe.Pointer(&header[0])
	_ = unsafe.Pointer(&miningConfig[0])
	_ = unsafe.Pointer(&proof[0])

	var out C.struct_CZKProof
	var errBuf [C.ERROR_MSG_MAX_SIZE]C.char

	buf := C.malloc(C.size_t(C.MAX_ZK_PROOF_SIZE))
	if buf == nil {
		return fmt.Errorf("failed to allocate proof buffer")
	}
	defer C.free(buf)

	out.proof_blob = (*C.uint8_t)(buf)

	cHeader := headerToC(header)

	rc := C.prove_plain_proof(
		&cHeader,
		(*[C.MINING_CONFIG_SERIALIZED_SIZE]C.uint8_t)(unsafe.Pointer(&miningConfig[0])),
		(*C.uint8_t)(unsafe.Pointer(&proof[0])),
		C.uintptr_t(len(proof)),
		&out,
		(*C.char)(unsafe.Pointer(&errBuf[0])),
	)

	if rc != 0 {
		return fmt.Errorf(C.GoString((*C.char)(unsafe.Pointer(&errBuf[0]))))
	}

	publicData := C.GoBytes(
		unsafe.Pointer(&out.public_data[0]),
		C.int(out.public_data_len),
	)

	proofData := C.GoBytes(
		unsafe.Pointer(out.proof_blob),
		C.int(out.proof_blob_len),
	)
         hashJackpot := C.GoBytes(
                unsafe.Pointer(&out.hash_jackpot[0]),
                32,
        )

        fmt.Printf("HASH JACKPOT = %x\n", hashJackpot)

	LastProof = &ZKProof{
		PublicData: publicData,
		ProofData:  proofData,
               HashJackpot: hashJackpot,
	}



	return nil
  
  
}
func VerifyWithNBits(header []byte, proof *ZKProof, nbits uint32) error {
	if proof == nil {
		return fmt.Errorf("nil proof")
	}

	var cProof C.struct_CZKProof

	// Public data
	cProof.public_data_len = C.uintptr_t(len(proof.PublicData))
	copy(
		(*[C.PUBLICDATA_MAX_SIZE]byte)(unsafe.Pointer(&cProof.public_data[0]))[:],
		proof.PublicData,
	)

	// Proof blob
	buf := C.malloc(C.size_t(len(proof.ProofData)))
	if buf == nil {
		return fmt.Errorf("malloc failed")
	}
	defer C.free(buf)

	C.memcpy(
		buf,
		unsafe.Pointer(&proof.ProofData[0]),
		C.size_t(len(proof.ProofData)),
	)

	cProof.proof_blob = (*C.uint8_t)(buf)
	cProof.proof_blob_len = C.uintptr_t(len(proof.ProofData))

	
var errBuf [C.ERROR_MSG_MAX_SIZE]C.char

cHeader := headerToC(header)

log.Println("C1")


rc := C.verify_zk_proof_v2_with_nbits(
    &cHeader,
    &cProof,
    C.uint32_t(nbits),
    (*C.char)(unsafe.Pointer(&errBuf[0])),
)


log.Println("C2")

if rc != 0 {
	return fmt.Errorf(C.GoString((*C.char)(unsafe.Pointer(&errBuf[0]))))
}

return nil

}


func VerifyNetwork(header []byte, proof *ZKProof,) error {
log.Println(">>> VerifyNetwork() called")

	if proof == nil {
		return fmt.Errorf("nil proof")
	}

	var cProof C.struct_CZKProof

	// Public data
	cProof.public_data_len = C.uintptr_t(len(proof.PublicData))
	copy(
		(*[C.PUBLICDATA_MAX_SIZE]byte)(unsafe.Pointer(&cProof.public_data[0]))[:],
		proof.PublicData,
	)

	// Proof blob
	buf := C.malloc(C.size_t(len(proof.ProofData)))
	if buf == nil {
		return fmt.Errorf("malloc failed")
	}
	defer C.free(buf)

	C.memcpy(
		buf,
		unsafe.Pointer(&proof.ProofData[0]),
		C.size_t(len(proof.ProofData)),
	)

	cProof.proof_blob = (*C.uint8_t)(buf)
	cProof.proof_blob_len = C.uintptr_t(len(proof.ProofData))

	
var errBuf [C.ERROR_MSG_MAX_SIZE]C.char

cHeader := headerToC(header)

rc := C.verify_zk_proof_v2(
    &cHeader,
    &cProof,
    (*C.char)(unsafe.Pointer(&errBuf[0])),
)

if rc != 0 {
	return fmt.Errorf(C.GoString((*C.char)(unsafe.Pointer(&errBuf[0]))))
}

return nil

}

