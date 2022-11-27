package domain

import (
	"fmt"
	"sync"

	"github.com/MickStanciu/work-in-progress/internal/util"
	"go.uber.org/zap"
)

// Factory a place where workers are picking up jobs
// and execute them
// - has a job queue
// - has a collection of workers
// - collects info about jobs that failed
// - collects info about jobs done
// - can stop the work
type Factory struct {
	logger      *zap.SugaredLogger
	jobCapacity int
	workers     int
	jobQueue    chan JobSvc

	mu       sync.Mutex
	jobsDone int
}

func (f *Factory) MarkJobDone() {
	f.mu.Lock()
	f.jobsDone++
	f.mu.Unlock()
}

func (f *Factory) CheckJobDone() int {
	f.mu.Lock()
	r := f.jobsDone
	f.mu.Unlock()
	return r
}

func (f *Factory) RegisterJob(j JobSvc) {
	fmt.Printf("Factory RegisterJob %s\n", j.GetID())
	f.jobQueue <- j
	if len(f.jobQueue) == f.jobCapacity {
		close(f.jobQueue)
	}
}

func (f *Factory) StartWork(doneChannel chan bool) {
	fmt.Println("Factory StartWork")
	// todo: PickJob treat error

	for job := range f.jobQueue {
		fmt.Printf("selecting job %s\n", job.GetID())
		go func(j JobSvc) {
			worker := Worker{}
			worker.PickJob(j)
			f.MarkJobDone()
		}(job)
	}

	//for {
	//	jobsDone := f.CheckJobDone()
	//	if jobsDone == f.jobCapacity {
	//		go func() {
	//			doneChannel <- true
	//		}()
	//		break
	//	}
	//}
	fmt.Println("FINISHED")
}

// FactoryConfig configuration for the Factory
type FactoryConfig struct {
	// how many jobs can be queued
	JobCapacity int
	// how many workers can work simultaneous in the factory
	Workers int
}

func NewFactory(cfg *FactoryConfig) *Factory {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger) // flushes buffer, if any

	return &Factory{
		logger:      util.GetLogger(),
		jobCapacity: cfg.JobCapacity,
		workers:     cfg.Workers,
		jobQueue:    make(chan JobSvc, cfg.JobCapacity),
	}
}
