package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	mu *sync.RWMutex
	data map[int]int
}

func NewCache() *Cache {
	return &Cache{
		mu: &sync.RWMutex{},
		data: make(map[int]int),
	}
}

var cache = NewCache()

func LongCalculation(n int) int {
	secondsToSleep := rand.Float64() * float64(n)
	time.Sleep(time.Duration(secondsToSleep))
	return n + 1
}

func CachedLongCalculation(n int) int {
	cache.mu.RLock()
	found, ok := cache.data[n]
	cache.mu.RUnlock()

	if !ok {
		value := LongCalculation(n)
		cache.mu.Lock()
		// Может нужно проверить существование еще раз
		cache.data[n] = value
		cache.mu.Unlock()
		return value
	}

	cache.mu.Unlock() // Лишний unlock

	return found
}

func main() {
	nums := []int{5, 10, 22}


	for _, n := range nums {
		go func() {
			val := CachedLongCalculation(n)
			fmt.Println(val)
		}()
	}
}