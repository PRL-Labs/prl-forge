package stratum

type Job struct {
	ID       string
	PrevHash string
	Coinb1   string
	Coinb2   string
	Merkle   []string
	Version  string
	NBits    string
	NTime    string
	Clean    bool
}

func NewDummyJob() *Job {
	return &Job{
		ID:       "1",
		PrevHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Coinb1:   "01000000",
		Coinb2:   "ffffffff",
		Merkle:   []string{},
		Version:  "20000000",
		NBits:    "1d00ffff",
		NTime:    "68555555",
		Clean:    true,
	}
}