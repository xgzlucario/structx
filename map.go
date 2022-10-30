package structx

import (
	"sync"
)

type Map[KEY, VALUE Value] map[KEY]VALUE

// The same as sync.Map
type SyncMap[K, V Value] struct {
	sync.RWMutex
	Map[K, V]
}

func NewMap[K, V Value]() Map[K, V] {
	return make(Map[K, V], 16)
}

func NewSyncMap[K, V Value]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		Map: NewMap[K, V](),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.Map[k] = v
}

func (m *SyncMap[K, V]) Load(k K) V {
	m.RLock()
	defer m.RUnlock()
	return m.Map[k]
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()
	delete(m.Map, key)
}

func (m *SyncMap[K, V]) Range(f func(k K, v V)) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.Map {
		f(k, v)
	}
}

func (m *SyncMap[K, V]) Len(key K) int {
	return len(m.Map)
}
