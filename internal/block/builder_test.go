package block

import "testing"

func TestBuildNilTemplate(t *testing.T) {
	_, err := Build(nil)
	if err == nil {
		t.Fatal("expected error for nil template")
	}
} // <-- ЛИПСВАШЕ ТАЗИ СКОБА

func TestBuildHeader(t *testing.T) {
	tpl := &Template{
		Version:      1,
		PreviousHash: make([]byte, 32),
		MerkleRoot:   make([]byte, 32),
		Timestamp:    1234567890,
		Bits:         0x1d00ffff,
	}

	h, err := Build(tpl)
	if err != nil {
		t.Fatalf("Build() failed: %v", err)
	}

	if h == nil {
		t.Fatal("header is nil")
	}

	raw, err := h.Serialize()
	if err != nil {
		t.Fatalf("Serialize() failed: %v", err)
	}

	if len(raw) != IncompleteHeaderSize+ProofCommitmentSize {
		t.Fatalf(
			"unexpected header size %d expected %d",
			len(raw),
			IncompleteHeaderSize+ProofCommitmentSize,
		)
	}
}
