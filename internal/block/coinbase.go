package block

import (
	
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/techobg/prl-forge/internal/tx"
)

type Coinbase struct {
	Height uint64
	Value  uint64
	Flags  string
}

func CoinbaseHash(cb Coinbase) ([]byte, error) {

	// -------- ScriptSig --------

	height := EncodeScriptNum(int64(cb.Height))

flags, err := hex.DecodeString(cb.Flags)
if err != nil {
	return nil, err
}

sig := make([]byte, 0)

sig = append(sig, byte(len(height)))
sig = append(sig, height...)

sig = append(sig, byte(len(flags)))
sig = append(sig, flags...)

script := tx.NewScript(sig)

	// -------- Coinbase input --------

	in := tx.NewCoinbaseInput(script)

	// -------- Output --------

out := tx.NewTxOut(
	int64(cb.Value),
	tx.NewScript([]byte{0x6a}), // OP_RETURN
)
	// -------- Transaction --------

	txn := tx.NewTransaction()

	txn.Version = 1
	txn.Inputs = []*tx.TxIn{in}
	txn.Outputs = []*tx.TxOut{out}
	txn.LockTime = 0

	raw, err := txn.Bytes()
	if err != nil {
		return nil, err
	}

	fmt.Println("========== COINBASE ==========")
	fmt.Printf("Raw TX : %x\n", raw)

	h1 := sha256.Sum256(raw)
	h2 := sha256.Sum256(h1[:])

	fmt.Printf("TXID   : %x\n", h2[:])
	fmt.Println("==============================")

	return h2[:], nil
}