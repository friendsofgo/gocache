package gocache

import "sync"

// Boundless is an in-memory cache implementation
// that stores all the given keys and values with
// no limits at all. Safe for concurrency.
type Boundless[K comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

// NewBoundless is a constructor method that
// initializes a new Boundless cache.
func NewBoundless[K comparable, V any]() *Boundless[K, V] {
	return &Boundless[K, V]{
		items: make(map[K]V),
	}
}

// Get returns the element stored for the given
// key or the zero value of the type V.
// Since the cache has no limits, it should
// return a value for every single key
// previously stored.
func (c *Boundless[K, V]) Get(key K) V {
	c.RLock()
	defer c.RUnlock()

	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	return item
}

// Set updates the element stored for the given
// key with the given value of the type V.
// Since the cache has no limits, it should
// never cause evictions.
func (c *Boundless[K, V]) Set(key K, value V) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}
