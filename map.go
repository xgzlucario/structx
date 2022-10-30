package structx

import (
	"sync"
)

type Map[K Value, V AnyValue] map[K]V

// The same as sync.Map
type SyncMap[K Value, V AnyValue] struct {
	sync.RWMutex
	Map[K, V]
}

// NewMap: new map
func NewMap[K Value, V AnyValue]() Map[K, V] {
	return make(Map[K, V], 16)
}

// NewSyncMap: new sync map
func NewSyncMap[K Value, V AnyValue]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		Map: NewMap[K, V](),
	}
}

// Store
func (m *SyncMap[K, V]) Store(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.Map[k] = v
}

// Load
func (m *SyncMap[K, V]) Load(k K) V {
	m.RLock()
	defer m.RUnlock()
	return m.Map[k]
}

// Delete
func (m *SyncMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()
	delete(m.Map, key)
}

// Range
func (m *SyncMap[K, V]) Range(f func(k K, v V)) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.Map {
		f(k, v)
	}
}

// Len
func (m *SyncMap[K, V]) Len(key K) int {
	return len(m.Map)
}
