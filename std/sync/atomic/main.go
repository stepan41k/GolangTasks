package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	x := atomic.Int64{}
	
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
			}()

			x.CompareAndSwap(x.Load(), int64(i))
			fmt.Println("Loaded: ", x.Load())
		}(i)
	}

	wg.Wait()

	fmt.Println("Result: ", x.Load())
}