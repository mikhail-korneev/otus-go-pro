package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type CacheElement struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	if c.queue.Len() == 0 {
		return nil, false
	}

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		value := item.Value.(CacheElement).value
		return value, true
	}

	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	e := CacheElement{key: key, value: value}

	if _, ok := c.items[e.key]; ok {
		item := c.items[e.key]
		item.Value = e
		c.queue.MoveToFront(item)
		c.items[e.key] = item
		return true
	}

	newItem := c.queue.PushFront(e)
	c.items[e.key] = newItem
	if c.queue.Len() > c.capacity {
		backItem := c.queue.Back()
		e := backItem.Value.(CacheElement)
		c.queue.Remove(backItem)
		delete(c.items, e.key)
	}
	return false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
