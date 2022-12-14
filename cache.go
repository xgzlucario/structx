package structx

import (
	"math"
	"sync/atomic"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	// duration of expired keys evictions
	GCDuration = time.Minute

	// duration of update current timestamp
	TickerDuration = time.Millisecond

	// default expiry time
	DefaultTTL = time.Minute * 10
)

const (
	NoTTL int64 = math.MaxInt64
)

type cacheItem[V any] struct {
	Value V
	Ttl   int64 // expiredTime
}

type Cache[K string, V any] struct {
	// current timestamp update by ticker
	_now int64

	// call when key-value expired
	onExpired cmap.RemoveCb[K, *cacheItem[V]]

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
		if item.Ttl > c.now() {
			return item.Value, true
		}
	}
	return
}

// Set
func (c *Cache[K, V]) Set(key K, value V, ttl ...time.Duration) {
	item := &cacheItem[V]{
		Value: value, Ttl: NoTTL,
	}
	// with ttl
	if len(ttl) > 0 {
		item.Ttl = c.now() + int64(ttl[0])
	}
	c.m.Set(key, item)
}

// MSet
func (c *Cache[K, V]) MSet(values map[K]V, ttl ...time.Duration) {
	items := make(map[K]*cacheItem[V], len(values))
	_ttl := NoTTL
	if len(ttl) > 0 {
		_ttl = int64(ttl[0])
	}
	// ttl
	for k, v := range values {
		items[k] = &cacheItem[V]{
			Value: v, Ttl: _ttl,
		}
	}
	c.m.MSet(items)
}

// Keys
func (c *Cache[K, V]) Keys() []K {
	return c.m.Keys()
}

// SetTTL
func (c *Cache[K, V]) SetTTL(key K, ttl time.Duration) bool {
	item, ok := c.m.Get(key)
	if ok {
		item.Ttl = c.now() + int64(ttl)
		return true
	}
	return false
}

// OnExpired
func (c *Cache[K, V]) OnExpired(f cmap.RemoveCb[K, *cacheItem[V]]) *Cache[K, V] {
	c.onExpired = f
	return c
}

// Remove
func (c *Cache[K, V]) Remove(key K) {
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
	for t := range c.m.IterBuffered() {
		if f(t.Key, t.Val.Value) {
			break
		}
	}
}

// RangeWithTTL
func (c *Cache[K, V]) RangeWithTTL(f func(key K, value V, ttl int64) bool) {
	for t := range c.m.IterBuffered() {
		if f(t.Key, t.Val.Value, t.Val.Ttl) {
			break
		}
	}
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

		for t := range c.m.IterBuffered() {
			// clear expired keys
			if t.Val.Ttl < c.now() {
				// onExpired
				if c.onExpired != nil {
					c.m.RemoveCb(t.Key, c.onExpired)
				} else {
					c.m.Remove(t.Key)
				}
			}
		}
	}
}

// Marshal
func (c *Cache[K, V]) MarshalJSON() ([]byte, error) {
	return c.m.MarshalJSON()
}

// Unmarshal
func (c *Cache[K, V]) UnmarshalJSON(src []byte) error {
	return c.m.UnmarshalJSON(src)
}

// Print
func (c *Cache[K, V]) Print() {
	c.m.Print()
}
