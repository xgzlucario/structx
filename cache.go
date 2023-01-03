package structx

import (
	"math"
	"sync/atomic"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
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

type Cache[K string, V any] struct {
	// current timestamp update by ticker
	_now int64

	// call when key-value expired
	onExpired cmap.RemoveCb[string, V]

	// data
	m *SyncMap[K, *cacheItem[V]]
}

func (c *Cache[K, V]) now() int64 {
	return atomic.LoadInt64(&c._now)
}

// NewCache
func NewCache[V any]() *Cache[string, V] {
	cache := &Cache[string, V]{
		m:    NewSyncMap[*cacheItem[V]](),
		_now: time.Now().UnixNano(),
	}

	go cache.eviction()
	go cache.ticker()

	return cache
}

// Get
func (c *Cache[K, V]) Get(key K) (v V, ok bool) {
	item, ok := c.m.Get(key)
	if ok {
		// check expired
		if item.ttl > c.now() {
			return item.value, true
		}
	}
	return
}

// Set
func (c *Cache[K, V]) Set(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[V]{
		value: value, ttl: NoTTL,
	}
	// with ttl
	if len(ttl) > 0 {
		item.ttl = c.now() + int64(ttl[0])
	}
	c.m.Set(key, item)
}

// MSet
func (c *Cache[K, V]) MSet(keys []K, values []V, ttl ...time.Duration) {
	items := make(map[K]*cacheItem[V], len(keys))
	_ttl := Expression(len(ttl) > 0, int64(ttl[0]), NoTTL)
	// ttl
	for i, v := range values {
		items[keys[i]] = &cacheItem[V]{
			value: v, ttl: _ttl,
		}
	}
	c.m.MSet(items)
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

// OnExpired
func (c *Cache[K, V]) OnExpired(f cmap.RemoveCb[string, V]) *Cache[K, V] {
	c.onExpired = f
	return c
}

// Delete
func (c *Cache[K, V]) Delete(key K) {
	c.m.Remove(key)
}

// Clear
func (c *Cache[K, V]) Clear() {
	c.m.Clear()
}

// Len
func (c *Cache[K, V]) Len() int {
	return c.m.Count()
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

// Scheduled update current timestamp
func (c *Cache[K, V]) ticker() {
	for c != nil {
		time.Sleep(TickerDuration)
		atomic.SwapInt64(&c._now, time.Now().UnixNano())
	}
}

// Scheduled expired keys evictions
func (c *Cache[K, V]) eviction() {
	for c != nil {
		time.Sleep(GCDuration)

		c.m.Range(func(key K, item *cacheItem[V]) bool {
			// clear expired keys
			if item.ttl < c.now() {
				// onExpired
				if c.onExpired != nil {
					// c.m.RemoveCb(key, c.onExpired)
				} else {
					c.m.Remove(key)
				}
			}
			return false
		})
	}
}

// Print
func (c *Cache[K, V]) Print() {
	c.m.Print()
}
