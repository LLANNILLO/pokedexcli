package pokecache

import (
	"time"
)

type CacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct {
	cacheEntries map[string]CacheEntry
	duration time.Duration
	mu *sync.Mutex
}

func NewCache(interval time.Duration) (Cache) {
	return Cache{
		cacheEntries: map[string]CacheEntry{}
		duration: dur
		mu &sync.Mutex{}
	}
}



