/**
	Implement a simple cache as a key-value pair
	Keys can by of any type, int, string, float...
**/

package utils

import (
	"sync"
)

// Cache is a simple Key-Value store
type Cache struct {
	items map[any] any
	lock  sync.Mutex
}

// NewCache creates a cache.
// One cache should be enough for an app, but not enforcing Singleton pattern.
func NewCache() *Cache {
	return &Cache{
		items: make(map[any]any),
	}
}

func (c *Cache) Set(key any, value any)() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key any) any {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, _ := c.items[key]
	return value
}

// Pop makes delete moot, just ignore the return value of Pop
func (c *Cache) Pop(key any) any {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.items[key]
	if ok {
		delete(c.items, key)
	}
	return value
}
