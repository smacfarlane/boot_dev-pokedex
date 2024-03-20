package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mu       *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		fmt.Println("tick")
		c.mu.Lock()

		for k, entry := range c.entries {
			fmt.Printf("%v\n", k)
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
