package gocache

import (
	"container/list"
	"sync"
)

// LRU is an in-memory cache implementation
// that uses the Least Recently Used (LRU)
// algorithm as replacement algorithm.
// Safe for concurrency.
type LRU[V any] struct {
	sync.RWMutex
	items map[string]*list.Element
	cap   int
	lru   *list.List
}

type entry[V any] struct {
	key string
	val V
}

// NewLRU is a constructor method that initializes
// a new LRU cache with a maximum capacity of cap.
func NewLRU[V any](cap int) *LRU[V] {
	return &LRU[V]{
		items: make(map[string]*list.Element),
		cap:   cap,
		lru:   list.New(),
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

	c.lru.MoveToFront(item)
	return item.Value.(*entry[V]).val
}

// Set updates the element stored for the given
// key with the given value of the type V.
// When the maximum capacity of the cache
// is reached, then it evicts cache entries by
// using the Least Recently Used (LRU) algorithm.
func (c *LRU[V]) Set(key string, val V) {
	c.Lock()
	defer c.Unlock()

	if item, ok := c.items[key]; ok {
		c.lru.MoveToFront(item)
		item.Value.(*entry[V]).val = val
		return
	}

	if len(c.items) == c.cap {
		c.evict()
	}

	newItem := &entry[V]{key: key, val: val}
	element := c.lru.PushFront(newItem)
	c.items[key] = element
}

func (c *LRU[V]) evict() {
	item := c.lru.Back()
	if item != nil {
		c.lru.Remove(item)
		kv := item.Value.(*entry[V])
		delete(c.items, kv.key)
	}
}
