package gocache

import "sync"

type CacheR[V any] struct {
	sync.RWMutex
	items map[string]*listItem[V]
	cap   int
	lru   *list[V]
}

func NewR[V any](cap int) *CacheR[V] {
	return &CacheR[V]{
		items: make(map[string]*listItem[V]),
		cap:   cap,
		lru:   new(list[V]),
	}
}

func (c *CacheR[V]) Get(key string) V {
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

func (c *CacheR[V]) Set(key string, val V) {
	c.Lock()
	defer c.Unlock()

	if len(c.items) == c.cap {
		c.evict()
	}

	newItem := &listItem[V]{key: key, val: val}
	c.lru.prepend(newItem)
	c.items[key] = newItem
}

func (c *CacheR[V]) evict() {
	evicted := c.lru.pop()

	if evicted != nil {
		delete(c.items, evicted.key)
	}
}
