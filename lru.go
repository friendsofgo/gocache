package gocache

import (
	"container/list"
	"sync"
)

// LRU is an in-memory cache implementation
// that uses the Least Recently Used (LRU)
// algorithm as replacement algorithm.
// Safe for concurrency.
type LRU[K comparable, V any] struct {
	sync.RWMutex
	items map[K]*list.Element
	cap   int
	lru   *list.List
}

type entry[K comparable, V any] struct {
	key K
	val V
}

// NewLRU is a constructor method that initializes
// a new LRU cache with a maximum capacity of cap.
func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	return &LRU[K, V]{
		items: make(map[K]*list.Element),
		cap:   cap,
		lru:   list.New(),
	}
}

// Get returns the element stored for the given
// key or the zero value of the type V.
func (c *LRU[K, V]) Get(key K) V {
	c.RLock()
	defer c.RUnlock()

	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	c.lru.MoveToFront(item)
	return item.Value.(*entry[K, V]).val
}

// Set updates the element stored for the given
// key with the given value of the type V.
// When the maximum capacity of the cache
// is reached, then it evicts cache entries by
// using the Least Recently Used (LRU) algorithm.
func (c *LRU[K, V]) Set(key K, val V) {
	c.Lock()
	defer c.Unlock()

	if item, ok := c.items[key]; ok {
		c.lru.MoveToFront(item)
		item.Value.(*entry[K, V]).val = val
		return
	}

	if len(c.items) == c.cap {
		c.evict()
	}

	newItem := &entry[K, V]{key: key, val: val}
	element := c.lru.PushFront(newItem)
	c.items[key] = element
}

func (c *LRU[K, V]) evict() {
	item := c.lru.Back()
	if item != nil {
		c.lru.Remove(item)
		kv := item.Value.(*entry[K, V])
		delete(c.items, kv.key)
	}
}
