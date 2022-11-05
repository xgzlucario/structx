package structx

import "time"

// public
var DefaultTTL = time.Minute * 10

type item[K comparable, V any] struct {
	key  K
	data V
	ttl  int64 // expirate time
}

type Cache[K comparable, V any] struct {
	m      *SyncMap[K, *item[K, V]]
	gcChan chan *item[K, V]
}

// NewCache
func NewCache[K comparable, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{
		m:      NewSyncMap[K, *item[K, V]](),
		gcChan: make(chan *item[K, V], 32),
	}
	// start gc
	cache.startGC()
	return cache
}

// Store
func (c *Cache[K, V]) Store(k K, v V, ttl ...time.Duration) {
	// set ttl
	if len(ttl) > 0 {
		value := &item[K, V]{
			key:  k,
			data: v,
			ttl:  time.Now().Add(ttl[0]).Unix(),
		}
		c.m.Store(k, value)
		c.gcChan <- value

	} else {
		c.m.Store(k, &item[K, V]{
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

func (c *Cache[K, V]) startGC() {
	gcList := NewList[*item[K, V]]()

	for {
		select {
		case value := <-c.gcChan:
			// sort with ttl
			for i, v := range gcList.Values {
				if value.ttl <= v.ttl {
					gcList.Insert(i, value)
					return
				}
			}
		default:
		}

		// expirate
		if value := gcList.Index(0); value != nil {
			if gcList.Index(0).ttl < time.Now().Unix() {
				gcList.LPop()
			}
		}
	}
}
