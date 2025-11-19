package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	numJobs     = 500
	workerCount = runtime.NumCPU()
)

func RunWorkerPool(ctx context.Context, jobs <-chan int, workerCount int) <-chan int {
	results := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-jobs:
					if !ok {
						return
					}

					select {
					case results <- val * val:
						fmt.Printf("Worker %d - job %d\n", i, val*val)
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	jobs := make(chan int)

	go func() {
		for i := 0; i < numJobs; i++ {
			jobs <- i
		}

		close(jobs)
	}()


	result := RunWorkerPool(ctx, jobs, workerCount)
	for res := range result {
		fmt.Println(res)
	}
}
