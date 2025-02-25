package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mutex sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entry: make(map[string]cacheEntry),
		mutex: sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.entry, key)
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		tickTime := <-ticker.C

		for key, entry := range c.entry {
			if tickTime.Sub(entry.createdAt) > interval {
				c.Delete(key)
			}
		}
	}
}
