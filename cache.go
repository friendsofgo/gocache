package gocache

type Cache struct {
	items map[string]interface{}
}

func New() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	item, ok := c.items[key]
	if !ok {
		return nil
	}

	return item
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}
