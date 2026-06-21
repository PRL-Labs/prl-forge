package stratum

import (
	"sync"
)

type JobManager struct {
	mu      sync.RWMutex
	jobs    map[string]*Job
	current *Job
}

func NewJobManager() *JobManager {
	return &JobManager{
		jobs: make(map[string]*Job),
	}
}

func (jm *JobManager) SetCurrent(job *Job) {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	jm.jobs[job.ID] = job
	jm.current = job
}

func (jm *JobManager) Current() *Job {
	jm.mu.RLock()
	defer jm.mu.RUnlock()

	return jm.current
}

func (jm *JobManager) Get(id string) (*Job, bool) {
	jm.mu.RLock()
	defer jm.mu.RUnlock()

	job, ok := jm.jobs[id]
	return job, ok
}