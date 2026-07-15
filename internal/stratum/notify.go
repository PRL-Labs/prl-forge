package stratum

import (
	"log"

	stratumadapter "github.com/techobg/prl-forge/internal/adapter/stratum"
)

var adapter = stratumadapter.NewSRBMinerAdapter()

func NotifyJob(session *Session, job *Job) error {
	if job == nil {
		log.Println("⚠️ No current job available")
		return nil
	}

	msg, err := adapter.Notify(job)
	if err != nil {
		return err
	}

	if err := session.Send(msg); err != nil {
		return err
	}

	log.Printf(
		"📦 Pearl job %s sent (height=%d)",
		job.ID,
		job.Height,
	)

	return nil
}

func SendCurrentJob(session *Session) error {
	job := CurrentJob()

	if job == nil {
		log.Println("❌ CurrentJob == nil")
		return nil
	}

	// Make per-session copy.
	jobCopy := *job

	jobCopy.Difficulty = session.Difficulty
	jobCopy.Target = job.Target


log.Printf(
    "DIFF=%.0f TARGET=%s",
    jobCopy.Difficulty,
    jobCopy.Target,
)


	return NotifyJob(session, &jobCopy)
}