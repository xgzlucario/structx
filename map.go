package structx

import "sync"

type Map[T Value] map[T]T

// The same as sync.Map
type SyncMap[T Value] struct {
	sync.RWMutex
	Map[T]
}

func (m *SyncMap[T]) Store(key, value T) {
	m.Lock()
	defer m.Unlock()
	m.Map[key] = value
}

func (m *SyncMap[T]) Load(key T) T {
	m.RLock()
	defer m.RUnlock()
	return m.Map[key]
}

func (m *SyncMap[T]) Delete(key T) {
	m.Lock()
	defer m.Unlock()
	delete(m.Map, key)
}

func (m *SyncMap[T]) Range(f func(k, v T)) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.Map {
		f(k, v)
	}
}

func (m *SyncMap[T]) Len(key T) int {
	return len(m.Map)
}
