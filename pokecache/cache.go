package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheData map[string]cacheEntry
	mu        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := &Cache{CacheData: map[string]cacheEntry{}, mu: sync.Mutex{}}
	go cache.reapLoop(duration)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	entry := cacheEntry{createdAt: time.Now(), val: val}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.CacheData[key] = entry
}

func (cache *Cache) Get(key string) (data []byte, available bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cacheEntry, available := cache.CacheData[key]
	if !available {
		return []byte{}, available
	}
	return cacheEntry.val, available
}

func (cache *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		<-ticker.C
		cache.mu.Lock()
		for k, v := range cache.CacheData {
			if time.Since(v.createdAt) > duration {
				delete(cache.CacheData, k)
			}
		}
		cache.mu.Unlock()
	}
}
