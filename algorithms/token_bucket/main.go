package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	ticker   *time.Ticker
	tokens   chan struct{}
	stopChan chan struct{}
	wg       *sync.WaitGroup
}

func NewTokenBucket(rate int, capacity int) *TokenBucket {
	tb := &TokenBucket{
		ticker:   time.NewTicker(time.Second / time.Duration(rate)),
		tokens:   make(chan struct{}, capacity),
		stopChan: make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}

	tb.wg.Add(1)
	go tb.tokenGenerator()

	return tb
}

func (tb *TokenBucket) tokenGenerator() {
	defer tb.wg.Done()

	for {
		select {
		case <-tb.stopChan:
			return
		case <-tb.ticker.C:
			select {
			case tb.tokens <- struct{}{}:
				// empty
			default:
				// empty
			}
		}
	}
}

func (tb *TokenBucket) Allow(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tb.tokens:
		return nil
	}
}

func (tb *TokenBucket) Stop() {
	close(tb.stopChan)
	tb.ticker.Stop()
	tb.wg.Wait()
}

func main() {
	limiter := NewTokenBucket(10, 10)
	defer limiter.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d ждет разрешения...\n", id)

			// Используем контекст, чтобы не ждать вечно.
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if err := limiter.Allow(ctx); err != nil {
				fmt.Printf("Горутина %d не получила разрешение: %v\n", id, err)
			} else {
				fmt.Printf("Горутина %d получила разрешение и выполняет работу\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершили работу.")
}
