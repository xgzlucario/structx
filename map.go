package structx

import (
	"fmt"

	"golang.org/x/exp/slices"
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
func (m Map[K, V]) Delete(key K) error {
	_, ok := m[key]
	if ok {
		delete(m, key)
		return nil
	}
	return errKeyNotFound(key)
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
	values := make([]V, 0, m.Len())
	for _, v := range m {
		values = append(values, v)
	}
	return slices.Clip(values)
}

// Len
func (m Map[K, V]) Len() int {
	return len(m)
}

// Marshal
func (m Map[K, V]) Marshal() ([]byte, error) {
	return marshalJSON(m)
}

// Unmarshal
func (m Map[K, V]) Unmarshal(src []byte) error {
	return unmarshalJSON(src, m)
}

// Print
func (m Map[K, V]) Print() {
	m.Range(func(k K, v V) bool {
		fmt.Printf("%+v -> %+v\n", k, v)
		return false
	})
}
