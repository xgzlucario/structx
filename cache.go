package structx

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var (
	GCDuration = time.Minute

	DefaultTTL       = time.Minute * 10
	NoTTL      int64 = math.MaxInt64
)

type cacheItem[V any] struct {
	value V
	ttl   int64 // expiredTime
}

type Cache[K Value, V any] struct {
	m   *SyncMap[K, *cacheItem[V]]
	now int64
}

// NewCache
func NewCache[K Value, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{
		m:   NewSyncMap[K, *cacheItem[V]](),
		now: time.Now().UnixNano(),
	}

	// start gc and ticter
	go cache.gabCollect()
	go cache.ticker()

	return cache
}

// Set
func (c *Cache[K, V]) Set(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[V]{
		value: value, ttl: NoTTL,
	}
	// with ttl
	if len(ttl) > 0 {
		item.ttl = c.now + int64(ttl[0])
	}
	c.m.Store(key, item)
}

// Sets
func (c *Cache[K, V]) Sets(keys []K, values []V) {
	items := make([]*cacheItem[V], len(keys))
	for i, v := range values {
		items[i] = &cacheItem[V]{
			value: v, ttl: NoTTL,
		}
	}
	c.m.StoreMany(keys, items)
}

// SetTTL
func (c *Cache[K, V]) SetTTL(key K, ttl time.Duration) bool {
	item, ok := c.m.Load(key)
	if ok {
		item.ttl = c.now + int64(ttl)
		return true
	}
	return false
}

// Load
func (c *Cache[K, V]) Load(key K) (v V, ok bool) {
	item, ok := c.m.Load(key)
	if ok {
		// expired
		if item.ttl < c.now {
			c.m.Delete(key)
			return
		}
		return item.value, true
	}
	return
}

// Delete
func (c *Cache[K, V]) Delete(key K) bool {
	return c.m.Delete(key)
}

// Clear: Clear data
func (c *Cache[K, V]) Clear() {
	c.m.Clear()
}

// Release: release cache object, set as nil
func (c *Cache[K, V]) Release() {
	c = nil
}

func (c *Cache[K, V]) Len() int {
	return c.m.Len()
}

func (c *Cache[K, V]) Range(f func(key K, value V) bool) {
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > c.now {
			return f(k, v.value)
		}
		return false
	})
}

func (c *Cache[K, V]) RangeWithTTL(f func(key K, value V, ttl int64) bool) {
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > c.now {
			return f(k, v.value, v.ttl)
		}
		return false
	})
}

func (c *Cache[K, V]) ticker() {
	for c != nil {
		time.Sleep(time.Millisecond)
		atomic.SwapInt64(&c.now, time.Now().UnixNano())
	}
}

func (c *Cache[K, V]) gabCollect() {
	for c != nil {
		time.Sleep(GCDuration)

		c.m.Lock()
		// clear expired keys
		for key, item := range c.m.m {
			if item.ttl < c.now {
				delete(c.m.m, key)
			}
		}
		c.m.Unlock()
	}
}

func (c *Cache[K, V]) Print() {
	fmt.Println("====== start ======")
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		fmt.Printf("%v -> %v expired(%v)\n", k, v.value, v.ttl < c.now)
		return false
	})
	fmt.Println("======= end =======")
}
