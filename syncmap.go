package structx

import "sync"

// SynMap: generic version of sync.Map
type SyncMap[K Value, V AnyValue] struct {
	sync.RWMutex
	m Map[K, V]
}

func NewSyncMap[K Value, V AnyValue]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: NewMap[K, V](),
	}
}

func (m *SyncMap[K, V]) Store(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *SyncMap[K, V]) Load(k K) V {
	m.RLock()
	defer m.RUnlock()
	return m.m[k]
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

func (m *SyncMap[K, V]) Len(key K) int {
	return len(m.m)
}
