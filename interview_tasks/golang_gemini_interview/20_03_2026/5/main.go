package main

import (
	"sync"
	"time"
)

// Механизм очистки (Cleanup): Просроченные записи не должны вечно висеть в памяти. Реализуй эффективный способ удаления старых записей (не обязательно удалять мгновенно в секунду истечения, но память должна освобождаться).
// Подумай о производительности: что будет, если ключей миллионы, а мы запустим очистку по всей мапе? (Реализуй хотя бы базовый подход, но проговори в комментариях, как бы ты это оптимизировал).

type Cache struct {
	cache           map[string]CacheItem
	mu              sync.RWMutex
	cleanupInterval time.Duration
	maxCleanupTime  time.Duration
}

type CacheItem struct {
	value interface{}
	ttl   time.Time
}

func NewCache(cleanupInterval time.Duration) *Cache {
	newCache := &Cache{
		cache:           make(map[string]CacheItem),
		cleanupInterval: cleanupInterval,
	}

	go func() {
		newCache.cleanup()
	}()

	return newCache
}

func (c *Cache) cleanup() {
	newTicker := time.NewTicker(c.cleanupInterval)
	defer newTicker.Stop()

	for range newTicker.C {
		curTime := time.Now()

		c.mu.Lock()
		defer c.mu.Unlock()

		for key, value := range c.cache {
			if time.Since(curTime) > c.maxCleanupTime {
				break
			}

			if time.Since(value.ttl) > 0 {
				delete(c.cache, key)
			}
		}
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheItem{value: value, ttl: time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, flag := c.cache[key]
	return val, flag
}
