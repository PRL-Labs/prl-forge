package block

import "bytes"

// BuildMerkleRoot изчислява Merkle Root от вече хеширани транзакции.
func BuildMerkleRoot(txHashes [][]byte) []byte {
	if len(txHashes) == 0 {
		return make([]byte, 32)
	}

	level := make([][]byte, len(txHashes))
	copy(level, txHashes)

	for len(level) > 1 {
		var next [][]byte

		for i := 0; i < len(level); i += 2 {
			left := level[i]
			right := left

			if i+1 < len(level) {
				right = level[i+1]
			}

			data := bytes.Join([][]byte{left, right}, nil)
			next = append(next, DoubleSHA256(data))
		}

		level = next
	}

	return level[0]
}
