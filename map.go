package structx

import (
	"fmt"
)

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

// Values
func (m Map[K, V]) Values() []V {
	values := make([]V, m.Len())
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
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
