// Требования к реализации:
// 1) Потокобезопасность: Метод Allow() будет вызываться одновременно из сотен параллельных горутин.
// 2) Алгоритм Скользящего окна: Если maxRequests = 5 и window = 1 * time.Second, то при вызове Allow() в момент времени T, метод должен проверить, сколько запросов было сделано в интервале [T−1s, T].
// 3) Если меньше 5 — зафиксировать текущий запрос, добавить его метку времени и вернуть true. Иначе — false.
// Управление памятью: Старые метки времени (выходящие за пределы window) должны регулярно подчищаться, чтобы память слайса не росла бесконечно!

package main

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	maxRequests int
	window      time.Duration
	timestamps []time.Time
	mu sync.Mutex
}

func NewLimiter(maxRequests int, window time.Duration) *Limiter {
	return &Limiter{
		maxRequests: maxRequests,
		window:      window,
		timestamps: make([]time.Time, 0, maxRequests),
		mu: sync.Mutex{},
	}
}

func (l *Limiter) Allow() bool {
	now := time.Now()
	
	l.mu.Lock()
	defer l.mu.Unlock()

	i := 0
	for len(l.timestamps) > 0 && now.Sub(l.timestamps[0]) > l.window {
		if time.Since(l.timestamps[0]) > l.window {
			i++
		} else {
			l.timestamps = l.timestamps[i:]
		}
	}
	
	if len(l.timestamps) == l.maxRequests {
		return false
	}

	l.timestamps = append(l.timestamps, time.Now())
	
	return true
}

func main() {
	limiter := NewLimiter(3, 500*time.Millisecond)

	for range 100 {
		time.Sleep(100 * time.Millisecond)
		go func() {
			fmt.Println(limiter.Allow())
		}()
	}
}
