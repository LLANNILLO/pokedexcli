package pokecache

import (
	"time"
	"sync"
)

type CacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	CacheEntries map[string]CacheEntry
	duration time.Duration
	mu *sync.RWMutex
}

func NewCache(interval time.Duration) (Cache) {
	return Cache{
		CacheEntries: map[string]CacheEntry{},
		duration: interval,
		mu: &sync.RWMutex{},
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheEntry := CacheEntry {
		createdAt: time.Now(),
		val: val,
	}

	c.CacheEntries[key] = newCacheEntry
}

func (c Cache) Get(key string) (result []byte, founded bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheEntry, ok := c.CacheEntries[key]
	
	if !ok {
		return nil, false
	}

	result = cacheEntry.val

	return result, true

}
