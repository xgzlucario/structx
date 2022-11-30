package structx

import (
	"fmt"
	"sync"
)

// ======================= Map =======================
type Map[K comparable, V any] map[K]V

// NewMap
func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

// Load
func (m Map[K, V]) Load(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

// Store
func (m Map[K, V]) Store(key K, value V) {
	m[key] = value
}

// StoreMany
func (m Map[K, V]) StoreMany(keys []K, values []V) {
	for i := range keys {
		m[keys[i]] = values[i]
	}
}

// Delete
func (m Map[K, V]) Delete(key K) bool {
	_, ok := m[key]
	if ok {
		delete(m, key)
		return true
	}
	return false
}

// Range
func (m Map[K, V]) Range(f func(K, V) bool) {
	for k, v := range m {
		if f(k, v) {
			return
		}
	}
}

// Len
func (m Map[K, V]) Len() int {
	return len(m)
}

// Print
func (m Map[K, V]) Print() {
	m.Range(func(k K, v V) bool {
		fmt.Printf("%+v -> %+v\n", k, v)
		return false
	})
}

// ======================= SyncMap =======================

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
	m.m.Store(key, value)
}

// StoreMany
func (m *SyncMap[K, V]) StoreMany(keys []K, values []V) {
	m.Lock()
	defer m.Unlock()
	m.m.StoreMany(keys, values)
}

// Load
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	return m.m.Load(key)
}

// Delete
func (m *SyncMap[K, V]) Delete(key K) bool {
	m.Lock()
	defer m.Unlock()
	return m.m.Delete(key)
}

// Range
func (m *SyncMap[K, V]) Range(f func(K, V) bool) {
	m.RLock()
	defer m.RUnlock()
	m.m.Range(f)
}

// Clear
func (m *SyncMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()
	m.m = NewMap[K, V]()
}

// Len
func (m *SyncMap[K, V]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return m.m.Len()
}

// Print
func (m *SyncMap[K, V]) Print() {
	m.RLock()
	defer m.RUnlock()
	m.m.Print()
}
