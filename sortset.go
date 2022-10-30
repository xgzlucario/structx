package structx

import (
	"fmt"
	"sort"
)

type elem[K Value, V Number] struct {
	Key   K
	Value V
}

type elems[K Value, V Number] []*elem[K, V]

type SortSet[K Value, V Number] struct {
	elems[K, V]
	Ascend bool
}

// NewSortSet: Return new sortset
func NewSortSet[K Value, V Number]() *SortSet[K, V] {
	return &SortSet[K, V]{
		elems:  make([]*elem[K, V], 0),
		Ascend: true,
	}
}

// Incr: Incr if key exist or create
func (s *SortSet[K, V]) Incr(key K, value V) {
	defer s.Sort()
	// exist
	if i := s.Index(key); i != nil {
		i.Value += value
		return
	}
	// not exist
	s.elems = append(s.elems, &elem[K, V]{
		Key:   key,
		Value: value,
	})
}

// Index: get item by key
func (s *SortSet[K, V]) Index(key K) *elem[K, V] {
	for _, v := range s.elems {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// Score: get key's score
func (s *SortSet[K, V]) Score(key K) (V, bool) {
	if i := s.Index(key); i != nil {
		return i.Value, true
	}
	var res V
	return res, false
}

// Keys: get all keys from sortset
func (s *SortSet[K, V]) Keys() []K {
	arr := make([]K, s.Len())
	for i, v := range s.elems {
		arr[i] = v.Key
	}
	return arr
}

func (s *SortSet[K, V]) Swap(i, j int) {
	s.elems[i], s.elems[j] = s.elems[j], s.elems[i]
}

func (s *SortSet[K, V]) Len() int {
	return len(s.elems)
}

func (s *SortSet[K, V]) Less(i, j int) bool {
	if s.Ascend {
		return s.elems[i].Value < s.elems[j].Value
	}
	return s.elems[i].Value > s.elems[j].Value
}

// Sort: use sort
func (s *SortSet[K, V]) Sort() {
	sort.Sort(s)
}

// Print: print values
func (s *SortSet[K, V]) Print() {
	fmt.Printf("sortset(%v)\n", s.Ascend)
	for _, i := range s.elems {
		fmt.Printf("%v => %v\n", i.Key, i.Value)
	}
}
