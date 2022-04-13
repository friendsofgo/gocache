package gocache

type Cache[V any] struct {
	items map[string]V
}

func New[V any]() *Cache[V] {
	return &Cache[V]{
		items: make(map[string]V),
	}
}

func (c *Cache[V]) Get(key string) V {
	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	return item
}

func (c *Cache[V]) Set(key string, value V) {
	c.items[key] = value
}
