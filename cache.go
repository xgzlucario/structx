package structx

import "time"

var DefaultTTL = time.Minute * 10

type cacheItem[K Value, V comparable] struct {
	key  K
	data V
	ttl  int64 // expireTime
}

type Cache[K Value, V comparable] struct {
	m      *SyncMap[K, *cacheItem[K, V]]
	gcChan chan *cacheItem[K, V]
	alive  bool
}

// NewCache
func NewCache[K Value, V comparable]() *Cache[K, V] {
	cache := &Cache[K, V]{
		m:      NewSyncMap[K, *cacheItem[K, V]](),
		gcChan: make(chan *cacheItem[K, V], 32),
		alive:  true,
	}
	// start gc
	go cache.startGC()
	return cache
}

// Store
func (c *Cache[K, V]) Store(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[K, V]{
		key:  key,
		data: value,
	}
	// with ttl
	if len(ttl) > 0 {
		item.ttl = time.Now().Add(ttl[0]).Unix()
		c.gcChan <- item
	}
	c.m.Store(key, item)
}

// Load
func (c *Cache[K, V]) Load(key K) (*V, bool) {
	value, ok := c.m.Load(key)
	if ok {
		return &value.data, ok
	} else {
		return nil, false
	}
}

// Clear
func (c *Cache[K, V]) Clear() {
	c.m.Clear()
}

// Release
func (c *Cache[K, V]) Release() {
	c.alive = false
}

func (c *Cache[K, V]) Len() int {
	return c.m.Len()
}

func (c *Cache[K, V]) startGC() {
	gcSet := NewZSet[K, int64]()

	for c.alive {
		select {
		case value := <-c.gcChan:
			gcSet.Set(value.key, value.ttl)
		default:
		}

		key, ttl, err := gcSet.GetByRank(0)
		// expired
		if err == nil && ttl < time.Now().Unix() {
			gcSet.Delete(key)
		}
	}
}
