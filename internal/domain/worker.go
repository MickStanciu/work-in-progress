package domain

import "fmt"

// Worker is a handler of jobs
type Worker struct {
}

func (w *Worker) PickJob(job JobSvc) error {
	fmt.Printf("Worker pick job %s\n", job.GetID())
	return job.Handle()
}
