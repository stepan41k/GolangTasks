package main

import (
	"fmt"
	"sync"
)

func merge(ch1, ch2 chan int) chan int {
	resChan := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		for v := range ch1 {
			resChan <- v
		}
		wg.Done()
	}()
	
	go func() {
		for v := range ch2 {
			resChan <- v
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(resChan)
	}()
	
	return resChan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i < 7; i+=2 {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 2; i < 7; i+=2 {
			ch2 <- i
		}
		close(ch2)
	}()

	for v := range merge(ch1, ch2) {
		fmt.Println(v)
	}
}