package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{}
	c.cacheEntries = make(map[string]cacheEntry)
	c.mu = &sync.Mutex{}
	go c.reaploop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	ce := cacheEntry{}
	ce.createdAt = now
	ce.val = val
	c.cacheEntries[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce, ok := c.cacheEntries[key]

	if !ok {
		return make([]byte, 0), false
	}

	return ce.val, true

}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	lastInterval := time.Now()
	defer ticker.Stop()

	for tick := range ticker.C {
		c.mu.Lock()
		for key, ce := range c.cacheEntries {
			if ce.createdAt.Before(lastInterval) {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
		lastInterval = tick
	}
}
