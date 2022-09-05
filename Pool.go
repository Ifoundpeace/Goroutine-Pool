package GoroutinePool

import (
	"sync"
	"time"
)

// Goroutine Executor

type GoroutinePool struct {
	MaximumWorkers int
	ActiveWorkers  int
	wg             sync.WaitGroup
}

func (executor *GoroutinePool) Submit(function func()) {
	for {
		if executor.ActiveWorkers < executor.MaximumWorkers {
			break
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
	executor.wg.Add(1)
	executor.ActiveWorkers++
	go func() {
		function()
		defer executor.wg.Done()
		executor.ActiveWorkers--
	}()
}

func (executor *GoroutinePool) Wait() {
	executor.wg.Wait()
}

// End

func CreateGoroutinePool(workers int) *GoroutinePool {
	return &GoroutinePool{
		workers,
		0, sync.WaitGroup{},
	}
}
