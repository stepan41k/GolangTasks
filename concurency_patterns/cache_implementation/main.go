package main

import (
	"fmt"
	"sync"
	"hash/fnv"
)

type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type Shard struct {
	data map[string]string
	mx sync.RWMutex
}

type MyCache struct {
	shards []*Shard
}

func NewMyCache(shardCount int64) *MyCache {
	shards := make([]*Shard, shardCount)

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}

	return &MyCache{
		shards: shards,
	}
}

func (c *MyCache) Set(k string, v string) {
	shard := c.getShard(k)

	shard.mx.Lock()
	defer shard.mx.Unlock()
	shard.data[k] = v
}

func (c *MyCache) Get(k string) (string, bool) {
	shard := c.getShard(k)

	shard.mx.RLock()
	defer shard.mx.RUnlock()

	value, ok := shard.data[k]
	return value, ok
}

func (c *MyCache) getShard(k string) *Shard {
	hasher := fnv.New32()
	_, _ = hasher.Write([]byte(k))
	hash := hasher.Sum32()

	return c.shards[hash % uint32(len(c.shards))]
}


func main() {
	cache := NewMyCache(10)
	cache.Set("salary", "500000")

	value, ok := cache.Get("salary")
	if !ok {
		fmt.Println("No results")
		return
	}

	fmt.Println("Значение:", value)
}