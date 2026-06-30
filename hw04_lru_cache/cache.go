package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value any
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(c.items[key])
		c.items[key].Value.(*cacheItem).value = value
	}

	if !ok {
		c.queue.PushFront(&cacheItem{key, value})
		c.items[key] = c.queue.Front()
	}

	if c.queue.Len() > c.capacity {
		delete(c.items, c.queue.Back().Value.(*cacheItem).key)
		c.queue.Remove(c.queue.Back())
	}

	return ok
}

func (c *lruCache) Get(key Key) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, ok
	}

	return nil, ok
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
