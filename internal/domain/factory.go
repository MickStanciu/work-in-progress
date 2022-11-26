package domain

import (
	"fmt"

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
}

func (f *Factory) RegisterJob(j JobSvc) {
	fmt.Printf("Factory RegisterJob %s\n", j.GetID())
}

func (f *Factory) StartWork(doneChannel chan bool) {
	fmt.Println("Factory StartWork")
	// todo: doneChannel
	// todo: PickJob treat error

	go func() {
		job := <-f.jobQueue
		worker := Worker{}
		worker.PickJob(job)
	}()
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
