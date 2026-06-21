package stratum

import (
	"github.com/techobg/prl-forge/internal/pool"
)

var (
	jobManager   = NewJobManager()
	shareManager = NewShareManager()
)

// Job е alias към pool.Job.
type Job = pool.Job

func SetCurrentJob(job *Job) {
	jobManager.SetCurrent(job)
}

func CurrentJob() *Job {
	return jobManager.Current()
}

func GetJob(id string) (*Job, bool) {
	return jobManager.Get(id)
}