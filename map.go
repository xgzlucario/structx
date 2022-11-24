package structx

import (
	"sync"
)

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

// SynMap: generic version of sync.Map
type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	m Map[K, V]
}

// NewSyncMap
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: NewMap[K, V](),
	}
}

// Store
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

// StoreMany
func (m *SyncMap[K, V]) StoreMany(keys []K, values []V) {
	if len(keys) != len(values) {
		panic("length is not equal between keys and values")
	}
	m.Lock()
	defer m.Unlock()
	for i := range keys {
		m.m[keys[i]] = values[i]
	}
}

// Load
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

// Load and delete
func (m *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.m[key]
	if ok {
		delete(m.m, key)
	}
	return v, ok
}

// Load or store
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	m.Lock()
	defer m.Unlock()
	v, ok := m.m[key]
	if !ok {
		m.m[key] = value
	}
	return v, ok
}

// Delete
func (m *SyncMap[K, V]) Delete(key K) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.m[key]
	if ok {
		delete(m.m, key)
		return true
	}
	return false
}

// Range
func (m *SyncMap[K, V]) Range(f func(K, V) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.m {
		if f(k, v) {
			return
		}
	}
}

// Clear map
func (m *SyncMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()
	m.m = NewMap[K, V]()
}

// Len
func (m *SyncMap[K, V]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}
