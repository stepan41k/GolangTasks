package main

import (
	"fmt"
	"sync"
)

// Условие: Напишите функцию для параллельного вычисления суммы квадратов чисел от 1 до n. Разделите работу на k горутин, каждая обрабатывает свой диапазон чисел. Пример: n = 10 → 385 (1² + 2² + ... + 10²)


func sumOfSquaresRange(start int, end int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int = 0
	for i := start; i <= end; i++ {
		sum += i * i
	}
	ch <- sum
}

func parallelSumOfSquares(n int, k int) int {
	if n < k {
		k = n
	}

	result := make(chan int, k)
	wg := &sync.WaitGroup{}

	step := n / k

	for i := 0; i < k; i++ {
		start := i * step + 1
		end := (i + 1) * step

		if i == k - 1 {
			end = n
		}
		
		wg.Add(1)

		go sumOfSquaresRange(start, end, result, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	totalSum := 0
	for v := range result {
		totalSum += v
	}

	return totalSum
}

func main() {
	n := 10
	k := 10

	result := parallelSumOfSquares(n, k)
	fmt.Println("result:", result)
}