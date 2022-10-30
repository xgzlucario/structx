package structx

import (
	"fmt"
	"sort"
)

type elem[K Value, V Number] struct {
	Key   K
	Value V
	Data  any
}

type elems[K Value, V Number] []*elem[K, V]

// ZSet: SortedSet
type ZSet[K Value, V Number] struct {
	elems[K, V]
}

// NewZSet: Return new ZSet
func NewZSet[K Value, V Number]() *ZSet[K, V] {
	return &ZSet[K, V]{
		elems: make([]*elem[K, V], 0, MAKE_SIZE),
	}
}

// Incr: Incr if key exist or create
func (s *ZSet[K, V]) Incr(key K, value V) {
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
func (s *ZSet[K, V]) Index(key K) *elem[K, V] {
	for _, v := range s.elems {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// Score: get key's score
func (s *ZSet[K, V]) Score(key K) (V, bool) {
	if i := s.Index(key); i != nil {
		return i.Value, true
	}
	var res V
	return res, false
}

// Rank: get key's rank
func (s *ZSet[K, V]) Rank(key K) int {
	for i, v := range s.elems {
		if v.Key == key {
			return i
		}
	}
	return -1
}

// Keys: get all keys
func (s *ZSet[K, V]) Keys() []K {
	arr := make([]K, s.Len())
	for i, v := range s.elems {
		arr[i] = v.Key
	}
	return arr
}

func (s *ZSet[K, V]) Swap(i, j int) {
	s.elems[i], s.elems[j] = s.elems[j], s.elems[i]
}

func (s *ZSet[K, V]) Len() int {
	return len(s.elems)
}

func (s *ZSet[K, V]) Less(i, j int) bool {
	return s.elems[i].Value < s.elems[j].Value
}

// Sort: use sort
func (s *ZSet[K, V]) Sort() {
	sort.Sort(s)
}

// Print: print values
func (s *ZSet[K, V]) Print() {
	fmt.Println("ZSet")
	for _, i := range s.elems {
		fmt.Printf("%v => %v\n", i.Key, i.Value)
	}
}

// Merge: merge two ZSets
func (this *ZSet[K, V]) Merge(s *ZSet[K, V]) {
	for _, i := range s.elems {
		this.Incr(i.Key, i.Value)
	}
	this.Sort()
}
