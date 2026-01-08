package main

import (
	"fmt"
)

// Создайте функцию, которая читает данные из канала и обрабатывает их пакетами (batches) по 5 элементов. Когда накапливается 5 элементов или канал закрывается, пакет обрабатывается.  Пример: input = (1,2,3,4,5,6,7,8) → batches ([1,2,3,4,5], [6,7,8])

func batchProcessor(input <-chan int, batchSize int) <-chan []int {
	output := make(chan []int)

	go func() {
		defer close(output)

		batch := make([]int, 0, batchSize)

		for v := range input {
			batch = append(batch, v)
			if len(batch) == batchSize {
				out := make([]int, batchSize)
				copy(out, batch)
				output <- out
				batch = batch[:0]
			}
		}

		if len(batch) > 0 {
			out := make([]int, len(batch))
			copy(out, batch)
			output <- out
		}
	}()

	return output
}

func main() {
	input := make(chan int)

    go func() {
        for v := range 8 {
            input <- v+1
        }
        close(input)
    }()

    for batch := range batchProcessor(input, 5) {
        fmt.Println(batch)
    }
}