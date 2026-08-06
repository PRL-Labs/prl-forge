//! FFI for converting an existing PlainProof into a ZK proof.

use std::os::raw::c_char;
use std::slice;

use zk_pow::api::proof::IncompleteBlockHeader;
use zk_pow::api::prove;
use zk_pow::ffi::plain_proof::PlainProof;

use crate::common::{
    acquire_cache,
    catch_panic,
    set_error_msg,
    CZKProof,
    MAX_ZK_PROOF_SIZE,
};

#[no_mangle]
pub unsafe extern "C" fn prove_plain_proof(
    _block_header: *const IncompleteBlockHeader,
    _mining_config: *const [u8; crate::common::MINING_CONFIG_SERIALIZED_SIZE],
    _proof_bytes: *const u8,
    _proof_len: usize,
    _zk_proof_out: *mut CZKProof,
    _error_msg_out: *mut c_char,
) -> i32 {
    let proof_bytes = unsafe {
    slice::from_raw_parts(_proof_bytes, _proof_len)
};

let proof = match PlainProof::deserialize_compat(proof_bytes) {
    Ok(p) => p,
    Err(e) => {
        set_error_msg(_error_msg_out, &format!("Invalid PlainProof: {}", e));
        return -1;
    }
};

println!("PROVE CACHE BEFORE");

let mut cache = acquire_cache();

println!("PROVE CACHE AFTER");

let result = match catch_panic(|| {
    println!("PROVE START");


println!("========== PROVE HEADER ==========");
println!("VERSION = {:08x}", (*_block_header).version);
println!("TIME    = {:08x}", (*_block_header).timestamp);
println!("NBITS   = {:08x}", (*_block_header).nbits);
println!("PREV    = {:02x?}", (*_block_header).prev_block);
println!("MRKL    = {:02x?}", (*_block_header).merkle_root);
println!("==================================");
    prove::zk_prove_plain_proof(
        *_block_header,
        &proof,
        &mut cache,
        false,
    )
}) {
    Ok(Ok(r)) => r,
    Ok(Err(e)) => {
        set_error_msg(_error_msg_out, &format!("Prove failed: {}", e));
        return -1;
    }
    Err(e) => {
        set_error_msg(_error_msg_out, &format!("Panic: {}", e));
        return -1;
    }
};

drop(cache);
println!("CACHE DROPPED");

println!("PROVE FINISHED");

let out = unsafe { &mut *_zk_proof_out };

out.public_data_len = result.public_data.len();
out.public_data[..result.public_data.len()]
    .copy_from_slice(&result.public_data);

let proof_blob = unsafe {
    slice::from_raw_parts_mut(out.proof_blob, MAX_ZK_PROOF_SIZE)
};

proof_blob[..result.proof_data.len()]
    .copy_from_slice(&result.proof_data);

out.proof_blob_len = result.proof_data.len();

out.hash_jackpot = result.hash_jackpot;

set_error_msg(_error_msg_out, "OK");

0
}