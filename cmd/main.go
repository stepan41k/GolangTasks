package main

import (
	"fmt"
	"sync"
)

var (
	countJobs = 5
	countWorkers = 3
)

func merge(jobs chan int, result chan int) {
	for v := range jobs {
		result <- v * v
	}
}

func main() {
	jobsChan := make(chan int)
	resultChan := make(chan int)

	go func() {
		for i := 1; i <= countJobs; i++ {
			jobsChan <- i
		}

		close(jobsChan)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(countWorkers)
	for i := 0; i < countWorkers; i++ {
		go func() {
			defer wg.Done()
			merge(jobsChan, resultChan)
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}