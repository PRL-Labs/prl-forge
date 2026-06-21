package stratum

import "log"

func NotifyJob(session *Session, job *Job) error {
	if job == nil {
		log.Println("⚠️ No current job available")
		return nil
	}

	// Serialize merkle properly (miners expect array, not Go slice)
	merkle := make([]string, len(job.Merkle))
	copy(merkle, job.Merkle)

	err := session.Notify(
		"mining.notify",
		[]any{
			job.ID,
			job.PrevHash,
			job.Coinb1,
			job.Coinb2,
			merkle,
			job.Version,
			job.NBits,
			job.NTime,
			job.Clean,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("📦 Job %s sent (height=%d)", job.ID, job.Height)
	return nil
}

func SendCurrentJob(session *Session) error {
	return NotifyJob(session, CurrentJob())
}