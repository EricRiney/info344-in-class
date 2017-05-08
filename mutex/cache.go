package main

import (
	"fmt"
	"sync"
	"time"
)

//CacheEntry represents an entry in the cache
type CacheEntry struct {
	value   string
	expires time.Time
}

//Cache represents a map[string]string that is safe
//for concurrent access
type Cache struct {
	mu      sync.RWMutex
	entries map[string]*CacheEntry
	quit    chan bool
}

//NewCache creates and returns a new Cache
func NewCache() *Cache {
	c := &Cache{
		entries: make(map[string]*CacheEntry),
		mu:      sync.RWMutex{},
		quit:    make(chan bool),
	}
	go c.startJanitor()
	return c
}

func (c *Cache) Close() {
	c.quit <- true
}

func (c *Cache) startJanitor() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			c.purgeExpired()
		case <-c.quit:
			return
		}
	}
}

func (c *Cache) purgeExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	nPurged := 0
	for key, entry := range c.entries {
		if now.After(entry.expires) {
			delete(c.entries, key)
			nPurged++
		}
	}
	fmt.Printf("purged %d entries\n", nPurged)
}

//Get returns the value associated with the requested key.
//The returned boolean will be false if the key was not
//in the cache.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry := c.entries[key]
	if entry == nil {
		return "", false
	}
	return entry.value, true
}

//Set sets the value associated with the given key.
//If the key is not yet in the cache, it will be added.
func (c *Cache) Set(key string, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := c.entries[key]
	if entry == nil {
		entry = &CacheEntry{}
		c.entries[key] = entry
	}
	entry.value = value
	entry.expires = time.Now().Add(ttl)
}
