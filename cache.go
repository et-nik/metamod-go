package metamod_go

import "sync"

type inmemoryCache[K comparable, V any] struct {
	cache map[K]V
	mu    sync.RWMutex
}

func newInmemoryCache[K comparable, V any]() *inmemoryCache[K, V] {
	return &inmemoryCache[K, V]{
		cache: make(map[K]V),
	}
}

func (c *inmemoryCache[K, V]) Get(s K) (value V, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.cache[s]
	if !ok {
		exists = false
		return
	}

	return v, true
}

func (c *inmemoryCache[K, V]) Set(s K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[s] = value
}
