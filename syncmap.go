package structx

import (
	"sync"

	"golang.org/x/exp/slices"
)

// SynMap: generic version of sync.Map
type SyncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  Map[K, V]
}

// NewSyncMap
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: NewMap[K, V](),
	}
}

// Get
func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Get(key)
}

// Gets
func (m *SyncMap[K, V]) Gets(keys []K) []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res := make([]V, 0, len(keys))
	for _, key := range keys {
		temp, ok := m.m.Get(key)
		if ok {
			res = append(res, temp)
		}
	}
	return slices.Clip(res)
}

// Set
func (m *SyncMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Set(key, value)
}

// Sets
func (m *SyncMap[K, V]) Sets(keys []K, values []V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := range keys {
		m.Set(keys[i], values[i])
	}
}

// Exist
func (m *SyncMap[K, V]) Exist(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Exist(key)
}

// Delete
func (m *SyncMap[K, V]) Delete(key K) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.m.Delete(key)
}

// Range
func (m *SyncMap[K, V]) Range(f func(K, V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.m.Range(f)
}

// Values
func (m *SyncMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Values()
}

// Clear
func (m *SyncMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m = NewMap[K, V]()
}

// Len
func (m *SyncMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Len()
}

// Marshal
func (m *SyncMap[K, V]) Marshal() ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return marshalJSON(m.m)
}

// Unmarshal
func (m *SyncMap[K, V]) Unmarshal(src []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return unmarshalJSON(src, m.m)
}

// Print
func (m *SyncMap[K, V]) Print() {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.m.Print()
}
