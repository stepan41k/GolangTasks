package main

import (
	"fmt"
	"sync"
)

func GoroutinesPrint(n int) {
    fmt.Printf("Goroitine %d\n", n)
}
 
func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			GoroutinesPrint(i)
		}(i)	
	}

	wg.Wait()
}