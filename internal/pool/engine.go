package pool

import "github.com/techobg/prl-forge/internal/pearl"

type Engine struct {
	Jobs    *JobManager
	Builder *Builder
}

func NewEngine() *Engine {
	return &Engine{
		Jobs:    NewJobManager(),
		Builder: NewBuilder(),
	}
}

func (e *Engine) CurrentJob() *Job {
	return e.Jobs.Current()
}

func (e *Engine) SetCurrentJob(job *Job) {
	e.Jobs.SetCurrent(job)
}

func (e *Engine) BuildJob(tpl *pearl.BlockTemplate) *Job {
	job := e.Builder.Build(tpl)
	e.SetCurrentJob(job)
	return job
}

func (e *Engine) NewJob() *Job {
	return e.Jobs.NewJob()
}

func (e *Engine) GetJob(id string) (*Job, bool) {
	return e.Jobs.Get(id)
}
