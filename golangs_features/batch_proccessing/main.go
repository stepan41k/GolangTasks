package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	batchSize = 10
	tickerPeriod = 1 * time.Second
	count int64 = 0
	countWorkers = 6
)

func proccessBatch(wg *sync.WaitGroup, batch []int) {
	// for i := 0; i < len(batch); i++ {
	// 	wg.Done()
	// }
	
	atomic.AddInt64(&count, int64(len(batch)))
	fmt.Printf("Обрабатывается батч размером %d\n",  len(batch))
}

func worker(wg *sync.WaitGroup, batch []int, ch chan int, ticker time.Ticker) {
	for {
		select {
		case item, ok := <-ch:
			if !ok {
				if len(batch) > 0 {
					proccessBatch(wg, batch)
					wg.Done()
				}
				return
			}
			batch = append(batch, item)
				
			if len(batch) == batchSize {
				proccessBatch(wg, batch)
				wg.Done()
				batch = nil
			}

			

		case <-ticker.C:
			if len(batch) > 0 {
				fmt.Println("timeout...")
				proccessBatch(wg, batch)
				wg.Done()
				batch = nil
			}	
		}	
	}
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	var batch []int
	ticker := time.NewTicker(tickerPeriod)
	defer ticker.Stop()

	for range countWorkers {
		go func() {
			worker(&wg, batch, ch, *ticker)
		}()
	}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}()
	}

	defer func() {
		wg.Wait()
		close(ch)
	}()

	time.Sleep(10* time.Second)

	fmt.Println(count, wg)
}