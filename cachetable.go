package structx

// CacheTable
type CacheTable[K Value, V any] struct {
	m *SyncMap[string, *Cache[K, V]]
}

// NewCacheTable
func NewCacheTable[K Value, V any]() *CacheTable[K, V] {
	table := &CacheTable[K, V]{
		m: NewSyncMap[string, *Cache[K, V]](),
	}
	return table
}

// Table: Return cache, auto create when table not exist
func (t *CacheTable[K, V]) Table(table string) *Cache[K, V] {
	cache, ok := t.m.Get(table)
	if ok {
		return cache
	}
	cache = NewCache[K, V]()
	t.m.Set(table, cache)
	return cache
}
