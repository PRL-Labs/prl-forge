package tx
type Transaction struct {
	Version int32

	// SegWit transaction
	SegWit bool

	Inputs  []*TxIn
	Outputs []*TxOut

	LockTime uint32
}

func NewTransaction() *Transaction {
	return &Transaction{
		Version:  1,
		Inputs:   make([]*TxIn, 0),
		Outputs:  make([]*TxOut, 0),
		LockTime: 0,
	}
}

func (tx *Transaction) AddInput(in *TxIn) {
	tx.Inputs = append(tx.Inputs, in)
}

func (tx *Transaction) AddOutput(out *TxOut) {
	tx.Outputs = append(tx.Outputs, out)
}

func (tx *Transaction) SerializeSize() int {
	size := 4 // Version

	size += VarIntSerializeSize(uint64(len(tx.Inputs)))
	for _, in := range tx.Inputs {
		size += in.SerializeSize()
	}

	size += VarIntSerializeSize(uint64(len(tx.Outputs)))
	for _, out := range tx.Outputs {
		size += out.SerializeSize()
	}

	size += 4 // LockTime

	return size
}
