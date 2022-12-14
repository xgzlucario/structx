package structx

import (
	"math/rand"
	"time"
)

/*
When the LSet length is less than LSET_MAX_SIZE, only use slice, otherwise use Map + List
*/
const LSET_MAX_SIZE = 32

/*
LSet (ListSet): map + list structure
LSet has richer api and faster Intersect, Union, Range operations than mapset
*/
type LSet[T comparable] struct {
	m Map[T, struct{}]
	*List[T]
}

// NewLSet: Create a new LSet
func NewLSet[T comparable](values ...T) *LSet[T] {
	ls := &LSet[T]{
		m:    NewMap[T, struct{}](),
		List: NewList[T](),
	}
	for _, v := range values {
		ls.Add(v)
	}
	return ls
}

// is enable to use map
func (s *LSet[T]) enable() bool {
	return s.Len() > LSET_MAX_SIZE || s.m.Len() > 0
}

// Add
func (s *LSet[T]) Add(key T) bool {
	if !s.Exist(key) {
		s.add(key)
		return true
	}
	return false
}

func (s *LSet[T]) add(key T) {
	s.RPush(key)
	// not use map
	if !s.enable() {
		return
	}

	if s.m.Len() > 0 {
		s.m[key] = struct{}{}
	} else {
		// init
		for _, v := range s.array {
			s.m[v] = struct{}{}
		}
	}
}

// Remove
func (s *LSet[T]) Remove(key T) bool {
	if s.Exist(key) {
		s.remove(key)
		return true
	}
	return false
}

func (s *LSet[T]) remove(key T) {
	// delete from map
	if s.enable() {
		delete(s.m, key)
	}
	// delete from list
	s.List.RemoveElem(key)
}

// Exist
func (s *LSet[T]) Exist(key T) bool {
	if s.Len() < LSET_MAX_SIZE {
		// slice
		for _, v := range s.array {
			if key == v {
				return true
			}
		}
	} else {
		// map
		_, ok := s.m[key]
		return ok
	}
	return false
}

// Copy
func (s *LSet[T]) Copy() *LSet[T] {
	lset := &LSet[T]{
		m:    make(Map[T, struct{}], s.Len()),
		List: NewList(s.array...),
	}
	// copy map
	if lset.enable() {
		for _, v := range s.array {
			lset.m[v] = struct{}{}
		}
	}
	return lset
}

// Equal: Compare members between two lsets is equal
func (s *LSet[T]) Equal(target *LSet[T]) bool {
	if s.Len() != target.Len() {
		return false
	}

	for _, val := range target.array {
		if !s.Exist(val) {
			return false
		}
	}
	return true
}

// Union: Return the union of two sets
func (s *LSet[T]) Union(t *LSet[T]) *LSet[T] {
	min, max := s.compareLength(t)
	// should copy max object
	max = max.Copy()

	for _, k := range min.array {
		max.Add(k)
	}
	return max
}

// Intersect
func (s *LSet[T]) Intersect(t *LSet[T]) *LSet[T] {
	min, max := s.compareLength(t)
	// should copy min object
	min = min.Copy()

	for _, k := range min.array {
		if !max.Exist(k) {
			min.remove(k)
		}
	}
	return min
}

// Difference
func (s *LSet[T]) Difference(t *LSet[T]) *LSet[T] {
	newS := NewLSet[T]()

	for _, key := range s.array {
		if !t.Exist(key) {
			newS.add(key)
		}
	}
	for _, key := range t.array {
		if !s.Exist(key) {
			newS.add(key)
		}
	}
	return newS
}

// IsSubSet
func (s *LSet[T]) IsSubSet(t *LSet[T]) bool {
	if t.Len() > s.Len() {
		return false
	}

	for _, v := range t.array {
		if !s.Exist(v) {
			return false
		}
	}
	return true
}

// LPop: Pop a elem from left
func (s *LSet[T]) LPop() T {
	v := s.List.LPop()
	if s.enable() {
		delete(s.m, v)
	}
	return v
}

// RPop: Pop a elem from right
func (s *LSet[T]) RPop() T {
	v := s.List.RPop()
	if s.enable() {
		delete(s.m, v)
	}
	return v
}

// RandomPop: Pop a random elem
func (s *LSet[T]) RandomPop() T {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(s.Len())

	s.Bottom(index)
	return s.RPop()
}

// Unmarshal: Unmarshal from JSON
func (s *LSet[T]) Unmarshal(src []byte) error {
	if err := unmarshalJSON(src, &s.array); err != nil {
		return err
	}
	*s = *NewLSet(s.array...)
	return nil
}

// Compare two lset length and return (*min, *max)
func (s1 *LSet[T]) compareLength(s2 *LSet[T]) (*LSet[T], *LSet[T]) {
	if s1.Len() < s2.Len() {
		return s1, s2
	}
	return s2, s1
}
