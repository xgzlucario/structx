package structx

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map/v2"
)

// SynMap: use ConcurrentMap
type SyncMap[K comparable, V any] struct {
	cmap.ConcurrentMap[K, V]
}

// NewSyncMap
func NewSyncMap[V any]() *SyncMap[string, V] {
	return &SyncMap[string, V]{
		cmap.New[V](),
	}
}

// NewSyncMapStringer
func NewSyncMapStringer[K cmap.Stringer, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		cmap.NewStringer[K, V](),
	}
}

// Range
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	for t := range m.IterBuffered() {
		if f(t.Key, t.Val) {
			break
		}
	}
}

// Print
func (m *SyncMap[K, V]) Print() {
	m.Range(func(k K, v V) bool {
		fmt.Printf("%+v -> %+v\n", k, v)
		return false
	})
	fmt.Println()
}
