package main

import (
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
	//TODO: protect this map with a RWMutex
	entries map[string]*CacheEntry
}

//NewCache creates and returns a new Cache
func NewCache() *Cache {
	c := &Cache{
		entries: make(map[string]*CacheEntry),
	}
	return c
}

//Get returns the value associated with the requested key.
//The returned boolean will be false if the key was not
//in the cache.
func (c *Cache) Get(key string) (string, bool) {
	//TODO: implement this method and
	//replace the return statement below
	return "", false
}

//Set sets the value associated with the given key.
//If the key is not yet in the cache, it will be added.
func (c *Cache) Set(key string, value string, ttl time.Duration) {
	//TODO: implement this method
}
