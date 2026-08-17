// По размеру: Как только накопилось batchSize событий (например, 10 штук).
// По таймауту: Если прошло время flushInterval (например, 500мс), а batchSize ещё не набрался, но в батче есть хотя бы 1 событие.
// При завершении работы (Close()): Все оставшиеся неотправленные события должны отправиться в flushFn, а фоновые горутины должны аккуратно завершиться (Graceful Shutdown).

// Требования:
// Потокобезопасность: Add будет вызываться одновременно из сотен параллельных горутин.
// Без утерь событий: При вызове Close() все накопленные, но ещё не отправленные события должны гарантированно записаться в flushFn.
// Без утечек горутин: Фоновая горутина тикера должна завершаться при вызове Close().
// Не отправлять пустые батчи: Если событий нет, flushFn вызывать не нужно.

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Event struct {
	ID   string
	Data string
}

type Batcher struct {
	batch         []Event
	closeChan     chan struct{}
	batchSize     int
	flushInterval time.Duration
	flushFn       func(events []Event)
	lastFlush     time.Time
	mu            sync.Mutex
}

func NewBatcher(batchSize int, flushInterval time.Duration, flushFn func(events []Event)) *Batcher {
	batcher := &Batcher{
		batch:         make([]Event, 0, batchSize),
		closeChan:     make(chan struct{}, 1),
		batchSize:     batchSize,
		flushInterval: flushInterval,
		flushFn:       flushFn,
		lastFlush:     time.Now(),
		mu:            sync.Mutex{},
	}

	go func() {
		ticker := time.NewTicker(flushInterval)
		defer ticker.Stop()

		for {
			select {
			case <-batcher.closeChan:
				return
			case <-ticker.C:
				now := time.Now()

				batcher.mu.Lock()
				var flushBatch []Event
				if now.Sub(batcher.lastFlush) > batcher.flushInterval && len(batcher.batch) > 0 {
					flushBatch = make([]Event, len(batcher.batch))
					copy(flushBatch, batcher.batch)
					batcher.batch = batcher.batch[:0]

				}
				batcher.mu.Unlock()

				if len(flushBatch) > 0 {
					batcher.flushFn(flushBatch)
				}
			}
		}
	}()

	return batcher
}

// Add добавляет новое событие в агрегатор.
func (b *Batcher) Add(event Event) {
	b.mu.Lock()

	b.batch = append(b.batch, event)

	if len(b.batch) == b.batchSize {
		b.lastFlush = time.Now()
		flushBatch := make([]Event, b.batchSize)
		copy(flushBatch, b.batch)
		b.batch = b.batch[:0]
		b.mu.Unlock()
		b.flushFn(flushBatch)
	} else {
		b.mu.Unlock()
	}
}

// Close корректно завершает работу: сбрасывает остатки событий и останавливает фоновую горутину.
func (b *Batcher) Close() {
	b.mu.Lock()

	if len(b.batch) > 0 {
		flushBatch := make([]Event, len(b.batch))
		copy(flushBatch, b.batch)
		b.batch = b.batch[:0]
		b.mu.Unlock()
		b.flushFn(flushBatch)
	} else {
		b.mu.Unlock()
	}

	b.closeChan <- struct{}{}
}

func main() {
	fn := func(events []Event) {
		for i, v := range events {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("process batch elem #:", i, v)
		}
	}

	batcher := NewBatcher(3, 1*time.Second, fn)

	wg := sync.WaitGroup{}
	for i := 1; i < 16; i++ {
		wg.Add(1)
		time.Sleep(200 * time.Millisecond)
		go func() {
			defer wg.Done()
			batcher.Add(Event{ID: strconv.Itoa(i), Data: "data" + strconv.Itoa(i)})
		}()
	}

	wg.Wait()
}
