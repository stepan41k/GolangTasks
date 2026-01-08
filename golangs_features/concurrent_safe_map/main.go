package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	data map[string]int
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data[key]
}
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}


func main() {
	newMap := &SafeMap{mu: sync.RWMutex{}, data: make(map[string]int)}
	word := "someword"
	wg := &sync.WaitGroup{}
	
	for i, v := range word {
		wg.Add(1)
		go func(k rune, v int) {
			defer wg.Done()
			newMap.Set(string(k), i)
		}(v, i)
	}

	for _, v := range word {
		wg.Add(1)
		go func(k rune) {
			defer wg.Done()
			fmt.Println(newMap.Get(string(v)))
		}(v)
	}

	wg.Wait()
}