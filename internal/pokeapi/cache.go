package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{cacheEntries: make(map[string]cacheEntry), mutex: &sync.Mutex{}}
	go cache.readLoop(time.Duration(interval))
	return cache
	
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.cacheEntries[key] = cacheEntry{val: val, createdAt: time.Now()}
	c.mutex.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.cacheEntries[key]
	if !ok {
		return []byte{}, false
	}
	return val.val, true
}

func (c Cache) readLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mutex.Lock()
		for key, value := range c.cacheEntries {
			if time.Since(value.createdAt) > interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}
