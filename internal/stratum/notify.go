package stratum

import "log"

func SendJob(session *Session, job *Job) error {

	err := session.Notify(
		"mining.notify",
		[]any{
			job.ID,
			job.PrevHash,
			job.Coinb1,
			job.Coinb2,
			job.Merkle,
			job.Version,
			job.NBits,
			job.NTime,
			job.Clean,
		},
	)

	if err != nil {
		return err
	}

	log.Printf("📦 Job %s sent", job.ID)

	return nil
}