package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func SimulateRequest(ctx context.Context) (int64, error) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ch := make(chan int64)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		counter.Add(1)
		ch <- counter.Load()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case count := <-ch:
		return count, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := SimulateRequest(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
