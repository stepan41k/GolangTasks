package main

import (
	"fmt"
	"sync"
)

// Одна горутина генерирует числа от 1 до n, несколько воркеров обрабатывают их (умножают на 2), одна горутина собирает результаты. Пример: n = 10, workers = 3 → (2, 4, 6, 8, 10, 12, 14, 16, 18, 20)

func fanOutFanIn(n int, workers int) []int {
	outChan := make(chan int)
	inChan := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			inChan <- i
		}

		close(inChan)
	}()
	
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()

			for v := range inChan {	
				outChan <- v * 2
			}	
		}()
	}

	go func() {
		wg.Wait()
		close(outChan)
	}()

	res := []int{}
	
	for v := range outChan {
		res = append(res, v)
		fmt.Println("value:", v)
	}
	

	return res
}

func main() {
	n := 10
	workers := 3

	fanOutFanIn(n, workers)
}	