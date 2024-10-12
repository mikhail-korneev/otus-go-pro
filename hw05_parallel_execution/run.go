package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// control channels
	work := make(chan Task, len(tasks))
	results := make(chan bool, n)
	stop := make(chan struct{})

	var wg sync.WaitGroup

	// run workers
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("worker %d: started\n", id)
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Printf("worker %d: stopped\n", id)
					return
				default:
					fmt.Printf("worker %d: working...\n", id)
					task, ok := <-work
					if !ok {
						fmt.Printf("worker %d: no more tasks, stopped\n", id)
						return
					}
					err := task()
					if err != nil {
						fmt.Printf("worker %d: error occured\n", id)
						results <- false
					} else {
						results <- true
					}
				}
			}
		}(i)
	}

	// send tasks to workers
	for _, task := range tasks {
		work <- task
	}
	close(work) // no more tasks

	// result checker
	success := 0
	failed := 0
	for i := 0; i < len(tasks); i++ {
		res := <-results
		if res {
			success++
		} else {
			failed++
		}
		fmt.Printf("current results: success=%d failed=%d\n", success, failed)
		if m > 0 && failed == m {
			fmt.Println("failed limit exceeded, stop!")
			break
		}
	}

	close(stop)
	fmt.Println("waiting for workers...")
	wg.Wait()
	fmt.Println("all workers stopped")
	close(results)
	fmt.Printf("final results: success=%d failed=%d\n", success, failed)

	if m > 0 && failed == m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
