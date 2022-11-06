package structx

import "time"

var DefaultTTL = time.Minute * 10

type cacheItem[K comparable, V any] struct {
	key  K
	data V
	ttl  int64 // expireTime
}

type Cache[K comparable, V any] struct {
	m      *SyncMap[K, *cacheItem[K, V]]
	gcChan chan *cacheItem[K, V]
	alive  bool
}

// NewCache
func NewCache[K comparable, V any]() *Cache[K, V] {
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
func (c *Cache[K, V]) Store(k K, v V, ttl ...time.Duration) {
	// with ttl
	if len(ttl) > 0 {
		value := &cacheItem[K, V]{
			key:  k,
			data: v,
			ttl:  time.Now().Add(ttl[0]).Unix(),
		}
		c.m.Store(k, value)
		c.gcChan <- value

	} else {
		c.m.Store(k, &cacheItem[K, V]{
			data: v,
		})
	}
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
			// sort with ttl
			gcSet.Set(value.key, value.ttl, nil)
		default:
		}

		if gcSet.Len() > 0 {
			// expire
			node := gcSet.GetDataByRank(0, true)
			if node.score < time.Now().Unix() {
				gcSet.Delete(node.key)
			}
		}
	}
}
