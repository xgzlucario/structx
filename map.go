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

// Get
func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

// Set
func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

// Exist
func (m Map[K, V]) Exist(key K) bool {
	_, ok := m[key]
	return ok
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

// Get
func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	return m.m.Get(key)
}

// Gets
func (m *SyncMap[K, V]) Gets(keys []K) []V {
	m.RLock()
	defer m.RUnlock()

	res := make([]V, 0, len(keys))
	for _, key := range keys {
		temp, ok := m.m.Get(key)
		if ok {
			res = append(res, temp)
		}
	}
	return res
}

// Set
func (m *SyncMap[K, V]) Set(key K, value V) {
	m.Lock()
	defer m.Unlock()
	m.m.Set(key, value)
}

// Sets
func (m *SyncMap[K, V]) Sets(keys []K, values []V) {
	m.Lock()
	defer m.Unlock()
	for i := range keys {
		m.Set(keys[i], values[i])
	}
}

// Exist
func (m *SyncMap[K, V]) Exist(key K) bool {
	m.RLock()
	defer m.RUnlock()
	return m.m.Exist(key)
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
