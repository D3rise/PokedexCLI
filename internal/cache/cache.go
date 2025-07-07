package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	key       string
	value     []byte
	createdAt time.Time
}

type Cache struct {
	mu       sync.RWMutex
	interval time.Duration
	cacheMap map[string]cacheEntry
}

func NewCache(d time.Duration) *Cache {
	cache := &Cache{cacheMap: make(map[string]cacheEntry)}

	go func() {
		ticker := time.NewTicker(d)

		for {
			<-ticker.C
			cache.reapLoop()
		}
	}()

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		key:       key,
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.cacheMap[key]
	return v.value, ok
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cacheMap {
		if time.Since(v.createdAt) > c.interval {
			delete(c.cacheMap, k)
		}
	}
}
