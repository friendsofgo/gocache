package gocache

import "sync"

type Boundless[V any] struct {
	sync.RWMutex
	items map[string]V
}

func NewBoundless[V any]() *Boundless[V] {
	return &Boundless[V]{
		items: make(map[string]V),
	}
}

func (c *Boundless[V]) Get(key string) V {
	c.RLock()
	defer c.RUnlock()

	var noop V

	item, ok := c.items[key]
	if !ok {
		return noop
	}

	return item
}

func (c *Boundless[V]) Set(key string, value V) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}
