package structx

// CacheTable
type CacheTable[V any] struct {
	m *SyncMap[string, *Cache[string, V]]
}

// NewCacheTable
func NewCacheTable[V any]() *CacheTable[V] {
	table := &CacheTable[V]{
		m: NewSyncMap[*Cache[string, V]](),
	}
	return table
}

// Table: Return cache, auto create when table not exist
func (t *CacheTable[V]) Table(table string) *Cache[string, V] {
	cache, ok := t.m.Get(table)
	if ok {
		return cache
	}
	cache = NewCache[V]()
	t.m.Set(table, cache)
	return cache
}

// DropTable
func (t *CacheTable[V]) DropTable(table string) {
	t.m.Remove(table)
}
