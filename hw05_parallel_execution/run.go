package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func isLimitExceeded(count, limit int) bool {
	return limit > 0 && count >= limit
}

func generator(tasks []Task, count *atomic.Int32, m int) <-chan Task {
	out := make(chan Task)
	go func() {
		defer close(out)
		for _, task := range tasks {
			if isLimitExceeded(int(count.Load()), m) {
				return
			}
			out <- task
		}
	}()
	return out
}

func worker(jobs <-chan Task, count *atomic.Int32) {
	for job := range jobs {
		if err := job(); err != nil {
			count.Add(1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	var count atomic.Int32
	jobs := generator(tasks, &count, m)

	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, &count)
		}()
	}

	wg.Wait()
	if isLimitExceeded(int(count.Load()), m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
