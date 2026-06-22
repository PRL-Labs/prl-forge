package stratum

import (
	"github.com/techobg/prl-forge/internal/pool"
)

// Job е alias към pool.Job.
type Job = pool.Job

var (
	engine       *pool.Engine
	shareManager = NewShareManager()
)

func SetEngine(e *pool.Engine) {
	engine = e
}

func CurrentJob() *Job {
	if engine == nil {
		return nil
	}

	return engine.CurrentJob()
}

func GetJob(id string) (*Job, bool) {
	if engine == nil {
		return nil, false
	}

	return engine.GetJob(id)
}