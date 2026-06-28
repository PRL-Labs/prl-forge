package stratumv2

type Job struct {
	ID string

	Height int64

	Header string

	Target string

	CertVersion int

	Clean bool
}
