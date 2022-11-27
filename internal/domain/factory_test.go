package domain_test

import (
	"fmt"
	"testing"

	"github.com/MickStanciu/work-in-progress/internal/domain"
)

func TestFactory_StartWork(t *testing.T) {
	cfg := &domain.FactoryConfig{
		JobCapacity: 10,
		Workers:     2,
	}
	factory := domain.NewFactory(cfg)
	doneCh := make(chan bool)

	for i := 0; i < 10; i++ {
		job := domain.Job{
			Id: fmt.Sprintf("ID_%d", i),
		}
		factory.RegisterJob(&job)
	}
	factory.StartWork(doneCh)

	<-doneCh
}
