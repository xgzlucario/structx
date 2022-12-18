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

// DropTable
func (t *CacheTable[K, V]) DropTable(table string) error {
	return t.m.Delete(table)
}

// Tables
func (t *CacheTable[K, V]) Tables() []string {
	tables := make([]string, 0, t.m.Len())
	t.m.Range(func(t string, _ *Cache[K, V]) bool {
		tables = append(tables, t)
		return false
	})
	return tables
}
