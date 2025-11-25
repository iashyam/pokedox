package pokecache

import (
	"sync"
	"time"
)

const sleepToCollectCache time.Duration = time.Hour

type Cache struct {
	val       []byte
	createdAt time.Time
}

type PokeCache struct {
	cacheEntries map[string]Cache
	mu           sync.RWMutex
}

func NewCache() *PokeCache {
	return &PokeCache{
		cacheEntries: make(map[string]Cache),
		mu:           sync.RWMutex{},
	}
}

func (cache *PokeCache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheEntries[key] = Cache{val: val, createdAt: time.Now()}
}

func (cache *PokeCache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	val, ok := cache.cacheEntries[key]
	return val.val, ok

}

func (cache *PokeCache) Reap(key string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	delete(cache.cacheEntries, key)
}

func (cache *PokeCache) ReapAll(interval time.Duration) {
	for {

		var keysToDelete []string
		cache.mu.RLock()
		for key, cse := range cache.cacheEntries {
			spentTime := time.Since(cse.createdAt)
			if spentTime > interval {
				keysToDelete = append(keysToDelete, key)
			}
		}
		cache.mu.RUnlock()

		cache.mu.Lock()
		for _, key := range keysToDelete {
			delete(cache.cacheEntries, key)
		}
		cache.mu.Unlock()

		time.Sleep(sleepToCollectCache)
	}
}
