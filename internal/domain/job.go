package domain

type JobSvc interface {
	GetID() string
	Handle() error
}

// Job todo
type Job struct {
	Id string
}

func (j *Job) GetID() string {
	return j.Id
}

func (j *Job) Handle() error {
	return nil
}
