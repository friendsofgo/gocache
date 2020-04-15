package gocache

type CacheR struct {
	items map[string]*listItem
	cap   int
	lru   *list
}

func NewR(cap int) *CacheR {
	return &CacheR{
		items: make(map[string]*listItem),
		cap:   cap,
		lru:   new(list),
	}
}

func (c *CacheR) Get(key string) interface{} {
	item, ok := c.items[key]
	if !ok {
		return nil
	}

	c.lru.update(item)
	return item.val
}

func (c *CacheR) Set(key string, val interface{}) {
	if len(c.items) == c.cap {
		c.evict()
	}

	newItem := &listItem{key: key, val: val}
	c.lru.prepend(newItem)
	c.items[key] = newItem
}

func (c *CacheR) evict() {
	evicted := c.lru.pop()

	if evicted != nil {
		delete(c.items, evicted.key)
	}
}
