package structx

import "fmt"

// HMap
// key -> field -> value
// (key, field) -> value
type HMap[K comparable, V any] struct {
	m Map[K, *Map[K, V]]
}

// NewHMap
func NewHMap[K comparable, V any]() *HMap[K, V] {
	return &HMap[K, V]{
		m: Map[K, *Map[K, V]]{},
	}
}

// Get
func (m *HMap[K, V]) Get(key, field K) (value V, ok bool) {
	n1, ok := m.m.Get(key)
	if ok {
		return n1.Get(field)
	}
	return
}

// Set
func (m *HMap[K, V]) Set(key, field K, value V) {
	node, ok := m.m.Get(key)
	if ok {
		node.Set(field, value)
		return
	}

	newMap := NewMap[K, V]()
	newMap.Set(field, value)
	m.m.Set(key, &newMap)
}

// Delete
func (m *HMap[K, V]) Delete(key K, field ...K) bool {
	if len(field) == 0 {
		return m.m.Delete(key)
	}

	n1, ok := m.m.Get(key)
	if ok {
		return n1.Delete(field[0])
	}
	return false
}

// Exist
func (m *HMap[K, V]) Exist(key, field K) bool {
	n1, ok := m.m.Get(key)
	if ok {
		return n1.Exist(key)
	}
	return false
}

// GetAllKeys
func (m *HMap[K, V]) GetAllKeys(key K) []K {
	n1, ok := m.m.Get(key)
	if ok {
		keys := make([]K, n1.Len())
		n1.Range(func(k K, _ V) bool {
			keys = append(keys, k)
			return false
		})
		return keys
	}
	return nil
}

// Len
func (m *HMap[K, V]) Len() int {
	return m.m.Len()
}

// Print
func (m *HMap[K, V]) Print() {
	m.m.Range(func(k1 K, m1 *Map[K, V]) bool {
		m1.Range(func(k2 K, v1 V) bool {
			fmt.Printf("%+v -> %+v -> %+v\n", k1, k2, v1)
			return false
		})
		return false
	})
}
