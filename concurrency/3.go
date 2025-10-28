package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	ch := make(chan bool, 1)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("отвисла")
		ch <- false
	}()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Произошел тик")
			ch <- true
		case value := <-ch:
			fmt.Println(value)
			wg.Wait()
			return
		}
	}
}