package pool

import (
	"fmt"
	"sync"
)

type JobManager struct {
	mu      sync.RWMutex
	jobs    map[string]*Job
	current *Job
	nextID  uint64
	maxJobs int
}

func NewJobManager() *JobManager {
	return &JobManager{
		jobs:    make(map[string]*Job),
		maxJobs: 32,
	}
}

func (jm *JobManager) NewJob() *Job {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	jm.nextID++

	job := &Job{
		ID:       fmt.Sprintf("%d", jm.nextID),
		PrevHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Coinb1:   "01000000",
		Coinb2:   "ffffffff",
		Merkle:   []string{},
		Version:  "20000000",
		NBits:    "1d00ffff",
		NTime:    "68555555",
		Clean:    true,
	}

	jm.jobs[job.ID] = job
	jm.current = job

	return job
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
	func (jm *JobManager) SetCurrent(job *Job) {
	jm.mu.Lock()
	defer jm.mu.Unlock()

	jm.jobs[job.ID] = job
	jm.current = job

}