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
	m    Map[T, struct{}]
	ls   *List[T]
	flag bool
}

// NewLSet: Create a new LSet
func NewLSet[T comparable](values ...T) *LSet[T] {
	ls := &LSet[T]{
		m:  NewMap[T, struct{}](),
		ls: NewList[T](),
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
	s.ls.RPush(key)
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
	if s.enable() {
		delete(s.m, key)
	}
	s.ls.Remove(key)
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

// Range
func (s *LSet[T]) Range(f func(T) bool) {
	s.ls.Range(func(i int, t T) bool {
		return f(t)
	})
}

// Copy
func (s *LSet[T]) Copy() *LSet[T] {
	lset := &LSet[T]{
		m:  make(Map[T, struct{}], s.Len()),
		ls: NewList(s.Members()...),
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

	for _, i := range s.Members() {
		var flag bool

		for _, j := range target.Members() {
			if i == j {
				flag = true
				break
			}
		}
		if !flag {
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
	min.Range(func(k T) bool {
		max.Add(k)
		return false
	})
	return max
}

func (this *LSet[T]) Intersect(t *LSet[T]) *LSet[T] {
	min, max := compareTwoLSet(this, t)
	// should copy min lset
	min = min.Copy()
	min.Range(func(k T) bool {
		if !max.Exist(k) {
			min.remove(k)
		}
		return false
	})
	return min
}

func (this *LSet[T]) Difference(t *LSet[T]) *LSet[T] {
	newSet := NewLSet[T]()
	this.Range(func(k T) bool {
		if !t.Exist(k) {
			newSet.add(k)
		}
		return false
	})
	t.Range(func(k T) bool {
		if !this.Exist(k) {
			newSet.add(k)
		}
		return false
	})
	return newSet
}

// LPop: Pop a elem from left
func (this *LSet[T]) LPop() (v T, ok bool) {
	v, ok = this.ls.LPop(), true
	if this.enable() {
		delete(this.m, v)
	}
	return
}

// RPop: Pop a elem from right
func (this *LSet[T]) RPop() (v T, ok bool) {
	v, ok = this.ls.RPop(), true
	if this.enable() {
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

	this.ls.Bottom(index)
	return this.RPop()
}

// Len
func (s *LSet[T]) Len() int {
	return s.ls.Len()
}

// Top: Move a elem to the top
func (s *LSet[T]) Top(elem T) bool {
	index := s.ls.Find(elem)
	if index < 0 {
		return false
	}
	s.ls.Top(index)
	return true
}

// Bottom: Move a elem to the bottom
func (s *LSet[T]) Bottom(elem T) bool {
	index := s.ls.Find(elem)
	if index < 0 {
		return false
	}
	s.ls.Bottom(index)
	return true
}

// Members: Get all members
func (s *LSet[T]) Members() array[T] {
	return s.ls.array
}

// Marshal: Marshal to bytes
func (s *LSet[T]) Marshal() ([]byte, error) {
	return s.ls.Marshal()
}

// Scan: Scan from bytes
func (s *LSet[T]) Scan(src []byte) error {
	var ls []T
	if err := sonic.Unmarshal(src, &s); err != nil {
		return err
	}
	*s = *NewLSet(ls...)
	return nil
}

// Compare two lset length and return (*min, *max)
func compareTwoLSet[T comparable](s1 *LSet[T], s2 *LSet[T]) (*LSet[T], *LSet[T]) {
	if s1.Len() <= s2.Len() {
		return s1, s2
	}
	return s2, s1
}

// Print
func (s *LSet[T]) Print() {
	s.ls.Print()
}
