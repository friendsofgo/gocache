package gocache

import "sync"

// LRU is an in-memory cache implementation
// that uses the Least Recently Used (LRU)
// algorithm as replacement algorithm.
// Safe for concurrency.
type LRU[V any] struct {
	sync.RWMutex
	items map[string]*listItem[V]
	cap   int
	lru   *list[V]
}

// NewLRU is a constructor method that initializes
// a new LRU cache with a maximum capacity of cap.
func NewLRU[V any](cap int) *LRU[V] {
	return &LRU[V]{
		items: make(map[string]*listItem[V]),
		cap:   cap,
		lru:   new(list[V]),
	}
}

// Get returns the element stored for the given
// key or the zero value of the type V.
func (c *LRU[V]) Get(key string) V {
	c.RLock()
	defer c.RUnlock()

	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	c.lru.update(item)

	return item.val
}

// Set updates the element stored for the given
// key with the given value of the type V.
// When the maximum capacity of the cache
// is reached, then it evicts cache entries by
// using the Least Recently Used (LRU) algorithm.
func (c *LRU[V]) Set(key string, val V) {
	c.Lock()
	defer c.Unlock()

	if len(c.items) == c.cap {
		c.evict()
	}

	newItem := &listItem[V]{key: key, val: val}
	c.lru.prepend(newItem)
	c.items[key] = newItem
}

func (c *LRU[V]) evict() {
	evicted := c.lru.pop()

	if evicted != nil {
		delete(c.items, evicted.key)
	}
}
