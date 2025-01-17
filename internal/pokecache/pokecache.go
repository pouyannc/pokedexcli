package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	fmt.Println("Added new cache entry!!!!!")
	cache.entries[key] = newEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	val, found := cache.entries[key]
	if !found {
		return nil, false
	}

	fmt.Println("Retrieving data from cache!!!!!!!!")
	return val.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		t := <-ticker.C
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
