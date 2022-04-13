package gocache

type CacheR[V any] struct {
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
	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	c.lru.update(item)

	return item.val
}

func (c *CacheR[V]) Set(key string, val V) {
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
