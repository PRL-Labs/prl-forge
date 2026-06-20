package pool

type Engine struct {
	Jobs *JobManager
}

func NewEngine() *Engine {
	return &Engine{
		Jobs: NewJobManager(),
	}
}

func (e *Engine) CurrentJob() *Job {
	return e.Jobs.Current()
}

func (e *Engine) NewJob() *Job {
	return e.Jobs.NewJob()
}

func (e *Engine) GetJob(id string) (*Job, bool) {
	return e.Jobs.Get(id)
}