package structx

import (
	"math/rand"
	"time"

	"github.com/bytedance/sonic"
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
	flag bool
	m    Map[T, struct{}]
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
	return s.Len() > LSET_MAX_SIZE || s.flag
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

	if s.flag {
		s.m[key] = struct{}{}
	} else {
		for _, v := range s.Members() {
			s.m[v] = struct{}{}
		}
		s.flag = true
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
	s.List.Remove(key)
}

// Exist
func (s *LSet[T]) Exist(key T) bool {
	if s.Len() < LSET_MAX_SIZE {
		// slice
		for _, v := range s.Members() {
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
		List: NewList(s.Members()...),
	}
	// copy map
	if lset.enable() {
		for _, v := range s.Members() {
			s.m[v] = struct{}{}
		}
	}
	return lset
}

// Equal: Compare members between two lsets is equal
func (s *LSet[T]) Equal(target *LSet[T]) bool {
	if s.Len() != target.Len() {
		return false
	}

	for _, val := range target.Members() {
		if !s.Exist(val) {
			return false
		}
	}
	return true
}

// Union: Return the union of two sets
func (this *LSet[T]) Union(t *LSet[T]) *LSet[T] {
	min, max := compareTwoLSet(this, t)
	// should copy max lset
	max = max.Copy()

	for _, k := range min.array {
		max.Add(k)
	}
	return max
}

// Intersect
func (this *LSet[T]) Intersect(t *LSet[T]) *LSet[T] {
	min, max := compareTwoLSet(this, t)
	// should copy min lset
	min = min.Copy()

	for _, k := range min.array {
		if !max.Exist(k) {
			min.remove(k)
		}
	}
	return min
}

// Difference
func (this *LSet[T]) Difference(t *LSet[T]) *LSet[T] {
	newSet := NewLSet[T]()

	for _, k := range this.array {
		if !t.Exist(k) {
			newSet.add(k)
		}
	}
	for _, k := range t.array {
		if !this.Exist(k) {
			newSet.add(k)
		}
	}
	return newSet
}

// IsSubSet
func (this *LSet[T]) IsSubSet(t *LSet[T]) bool {
	if t.Len() > this.Len() {
		return false
	}

	for _, v := range t.array {
		if !this.Exist(v) {
			return false
		}
	}
	return true
}

// LPop: Pop a elem from left
func (this *LSet[T]) LPop() (v T, ok bool) {
	v, ok = this.List.LPop()
	if ok && this.enable() {
		delete(this.m, v)
	}
	return
}

// RPop: Pop a elem from right
func (this *LSet[T]) RPop() (v T, ok bool) {
	v, ok = this.List.RPop()
	if ok && this.enable() {
		delete(this.m, v)
	}
	return
}

// RandomPop: Pop a random elem
func (this *LSet[T]) RandomPop() (v T, ok bool) {
	if this.Len() == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(this.Len())

	this.Bottom(index)
	return this.RPop()
}

// Members: Get all members
func (s *LSet[T]) Members() array[T] {
	return s.array
}

// Scan: Scan from bytes
func (s *LSet[T]) ScanJSON(src []byte) error {
	if err := sonic.Unmarshal(src, &s.array); err != nil {
		return err
	}
	*s = *NewLSet(s.array...)
	return nil
}

// Compare two lset length and return (*min, *max)
func compareTwoLSet[T comparable](s1 *LSet[T], s2 *LSet[T]) (*LSet[T], *LSet[T]) {
	if s1.Len() <= s2.Len() {
		return s1, s2
	}
	return s2, s1
}
