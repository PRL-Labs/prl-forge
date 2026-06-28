//go:build ignore

package zkpow

/*
#cgo CFLAGS: -I${SRCDIR}/../../ffi/include
#cgo LDFLAGS: -L${SRCDIR}/../../ffi/lib -lzk_pow_ffi


#include "zk_pow_ffi.h"
*/
import "C"

import "unsafe"

func verify(header, proof []byte) int {
	if len(header) == 0 || len(proof) == 0 {
		return -1
	}

	return int(C.verify_zk_proof_v2(
		(*C.uint8_t)(unsafe.Pointer(&header[0])),
		C.uint32_t(len(header)),
		(*C.uint8_t)(unsafe.Pointer(&proof[0])),
		C.uint32_t(len(proof)),
	))
}
