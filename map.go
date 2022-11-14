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

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: NewMap[K, V](),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *SyncMap[K, V]) Load(k K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, key)
}

func (m *SyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.m {
		if f(k, v) {
			return
		}
	}
}

func (m *SyncMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()
	m.m = NewMap[K, V]()
}

func (m *SyncMap[K, V]) Len() int {
	return len(m.m)
}
