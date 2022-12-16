package structx

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

// Modifiable configuration
var (
	// duration of expired keys evictions
	GCDuration = time.Minute

	// duration of update current timestamp
	TickerDuration = time.Millisecond

	// default expiry time
	DefaultTTL = time.Minute * 10
)

// Non-modifiable configuration
const (
	NoTTL int64 = math.MaxInt64
)

type cacheItem[V any] struct {
	value V
	ttl   int64 // expiredTime
}

type Cache[K Value, V any] struct {
	// current timestamp update by ticker
	_now int64

	// call when key-value expired
	onExpired func(K, V)

	// data
	m *SyncMap[K, *cacheItem[V]]
}

func (c *Cache[K, V]) now() int64 {
	return atomic.LoadInt64(&c._now)
}

// NewCache
func NewCache[K Value, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{
		m:    NewSyncMap[K, *cacheItem[V]](),
		_now: time.Now().UnixNano(),
	}

	// start gc and ticker
	go cache.gabCollect()
	go cache.ticker()

	return cache
}

// Store
func (c *Cache[K, V]) Store(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[V]{
		value: value, ttl: NoTTL,
	}
	// with ttl
	if len(ttl) > 0 {
		item.ttl = c.now() + int64(ttl[0])
	}
	c.m.Set(key, item)
}

// StoreMany
func (c *Cache[K, V]) StoreMany(keys []K, values []V, ttl ...time.Duration) {
	items := make([]*cacheItem[V], len(keys))
	// ttl
	_ttl := NoTTL
	if len(ttl) > 0 {
		_ttl = int64(ttl[0])
	}

	for i, v := range values {
		items[i] = &cacheItem[V]{
			value: v, ttl: _ttl,
		}
	}
	c.m.Sets(keys, items)
}

// SetTTL
func (c *Cache[K, V]) SetTTL(key K, ttl time.Duration) bool {
	item, ok := c.m.Get(key)
	if ok {
		item.ttl = c.now() + int64(ttl)
		return true
	}
	return false
}

// Load
func (c *Cache[K, V]) Load(key K) (v V, ok bool) {
	item, ok := c.m.Get(key)
	if ok {
		// check expired
		if item.ttl > c.now() {
			return item.value, true
		}
	}
	return
}

// OnExpired
func (c *Cache[K, V]) OnExpired(f func(K, V)) *Cache[K, V] {
	c.onExpired = f
	return c
}

// Delete
func (c *Cache[K, V]) Delete(key K) bool {
	return c.m.Delete(key)
}

// Clear
func (c *Cache[K, V]) Clear() {
	c.m.Clear()
}

// Len
func (c *Cache[K, V]) Len() int {
	return c.m.Len()
}

// Range
func (c *Cache[K, V]) Range(f func(key K, value V) bool) {
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > c.now() {
			return f(k, v.value)
		}
		return false
	})
}

// RangeWithTTL
func (c *Cache[K, V]) RangeWithTTL(f func(key K, value V, ttl int64) bool) {
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		if v.ttl > c.now() {
			return f(k, v.value, v.ttl)
		}
		return false
	})
}

func (c *Cache[K, V]) ticker() {
	for c != nil {
		time.Sleep(TickerDuration)
		atomic.SwapInt64(&c._now, time.Now().UnixNano())
	}
}

func (c *Cache[K, V]) gabCollect() {
	for c != nil {
		time.Sleep(GCDuration)
		c.m.Lock()
		// clear expired keys
		for key, item := range c.m.m {
			if item.ttl < c.now() {
				// onExpired
				if c.onExpired != nil {
					c.onExpired(key, item.value)
				}
				delete(c.m.m, key)
			}
		}
		c.m.Unlock()
	}
}

// Print
func (c *Cache[K, V]) Print() {
	c.m.Range(func(k K, v *cacheItem[V]) bool {
		fmt.Printf("%+v -> %+v\n", k, v.value)
		return false
	})
}
