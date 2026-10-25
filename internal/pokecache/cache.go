package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	CacheEntries map[string]cacheEntry
	duration time.Duration
	mu *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		CacheEntries: map[string]cacheEntry{},
		duration: interval,
		mu: &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheEntry := cacheEntry {
		createdAt: time.Now().UTC(),
		val: val,
	}

	c.CacheEntries[key] = newCacheEntry
}

func (c *Cache) Get(key string) (result []byte, founded bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.CacheEntries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.CacheEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.CacheEntries, k)
		}
	}
}
