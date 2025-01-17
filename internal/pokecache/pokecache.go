package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu      *sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:      &sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.entries[key] = newEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	val, found := cache.entries[key]
	return val.val, found
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for t := range ticker.C {
		cache.mu.Lock()
		fmt.Println("REAPING CACHE!!!!!!")
		for k, v := range cache.entries {
			if t.After(v.createdAt) {
				delete(cache.entries, k)
			}
		}
		cache.mu.Unlock()
	}
}
