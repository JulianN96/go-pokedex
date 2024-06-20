package pokecache

import (
	"sync"
	"time"
)

//Create a struct for the cache
type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

//Create a struct for each entry into the cache
type cacheEntry struct{
	val []byte
	createdAt time.Time
}

//Initialise the cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

//Generic function to add things to cache
func (c *Cache) Add(key string, val[]byte){
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}
}

//Generic function to get things from cache
func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

//Function to call the delete cache functiong every X minutes
func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{
		//this will run every interval
		c.reap(interval)
	}
}
//Function to delete items from cache
func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k,v := range c.cache{
		if v.createdAt.Before(timeAgo){
			delete(c.cache, k)
		}
	}
}
