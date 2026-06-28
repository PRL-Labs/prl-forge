package pool

import "log"

func SubmitBlock(job *Job) error {

	log.Printf(
    "🚀 SubmitBlock: job=%s height=%d proof=%d hs=%d template=%v",
    job.ID,
    job.Height,
    len(job.Proof),
    job.HS,
    job.Template != nil,
)

	return nil
}