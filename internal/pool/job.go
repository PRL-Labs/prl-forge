package pool

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