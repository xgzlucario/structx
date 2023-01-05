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

// Print
func (m *SyncMap[K, V]) Print() {
	for t := range m.IterBuffered() {
		fmt.Printf("%+v -> %+v\n", t.Key, t.Val)
	}
	fmt.Println()
}
