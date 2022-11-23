package structx

import (
	"math"
	"time"
)

var GCDuration = time.Minute

type cacheItem[V any] struct {
	data V
	ttl  int64 // expiredTime
}

type Cache[K Value, V any] struct {
	m     *SyncMap[K, *cacheItem[V]]
	alive bool
}

// NewCache
func NewCache[K Value, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{
		m:     NewSyncMap[K, *cacheItem[V]](),
		alive: true,
	}
	// start gc
	go cache.startGC()
	return cache
}

// Store
func (c *Cache[K, V]) Store(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[V]{
		data: value,
		ttl:  math.MaxInt64,
	}
	// with ttl
	if len(ttl) > 0 {
		item.ttl = time.Now().Add(ttl[0]).UnixNano()
	}
	c.m.Store(key, item)
}

// Load
func (c *Cache[K, V]) Load(key K) (v V, ok bool) {
	item, ok := c.m.Load(key)
	if ok {
		// expired
		if item.ttl < time.Now().UnixNano() {
			c.m.Delete(key)
			return
		}
		return item.data, true
	}
	return
}

// Delete
func (c *Cache[K, V]) Delete(key K) bool {
	return c.m.Delete(key)
}

// Clear
func (c *Cache[K, V]) Clear() {
	c.m.Clear()
}

// Release
func (c *Cache[K, V]) Release() {
	c.alive = false
	c = nil
}

func (c *Cache[K, V]) Len() int {
	return c.m.Len()
}

func (c *Cache[K, V]) Range(f func(key K, value V) bool) {
	now := time.Now().UnixNano()
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > now {
			return f(k, v.data)
		}
		return false
	})
}

func (c *Cache[K, V]) RangeWithTTL(f func(key K, value V, ttl int64) bool) {
	now := time.Now().UnixNano()
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > now {
			return f(k, v.data, v.ttl)
		}
		return false
	})
}

func (c *Cache[K, V]) startGC() {
	for c != nil && c.alive {
		time.Sleep(GCDuration)

		c.m.Lock()
		now := time.Now().UnixNano()
		// clear expired keys
		for key, item := range c.m.m {
			if item.ttl > now {
				delete(c.m.m, key)
			}
		}
		c.m.Unlock()
	}
}
