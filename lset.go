package structx

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bytedance/sonic"
)

/*
LSet (ListSet): map + list structure
LSet has richer api and faster Intersect, Union, Range operations than mapset
*/
type LSet[T comparable] struct {
	m  Map[T, struct{}]
	ls *List[T]
}

// NewLSet: Create a new LSet from list
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

// Add
func (s *LSet[T]) Add(key T) bool {
	_, ok := s.m[key]
	if ok {
		return false
	}
	s.add(key)
	return true
}

func (s *LSet[T]) add(key T) {
	s.ls.RPush(key)
	s.m[key] = struct{}{}
}

// Remove
func (s *LSet[T]) Remove(key T) bool {
	_, ok := s.m[key]
	if ok {
		s.remove(key)
		return true
	}
	return false
}

func (s *LSet[T]) remove(key T) {
	delete(s.m, key)
	s.ls.RemoveElem(key)
}

// Exist
func (s *LSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Range
func (s *LSet[T]) Range(f func(k T) bool) {
	for _, v := range s.ls.Array {
		if f(v) {
			return
		}
	}
}

// Copy
func (s *LSet[T]) Copy() *LSet[T] {
	newLSet := &LSet[T]{
		m:  make(Map[T, struct{}], s.Len()),
		ls: NewList(s.Members()...),
	}
	// copy map
	for _, v := range s.Members() {
		s.m[v] = struct{}{}
	}
	return newLSet
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
	if this.Len() > 0 {
		v = this.ls.LPop()
		delete(this.m, v)
		ok = true
	}
	return
}

// RPop: Pop a elem from right
func (this *LSet[T]) RPop() (v T, ok bool) {
	if this.Len() > 0 {
		v = this.ls.RPop()
		delete(this.m, v)
		ok = true
	}
	return
}

// RandomPop: Pop a random elem
func (this *LSet[T]) RandomPop() (v T, ok bool) {
	if this.Len() > 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(this.Len())

		this.ls.Bottom(index)
		return this.RPop()
	}
	return
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
func (s *LSet[T]) Members() Array[T] {
	return s.ls.Array
}

// Marshal: Marshal to bytes
func (s *LSet[T]) Marshal() ([]byte, error) {
	return sonic.Marshal(s.Members())
}

// Scan: Scan from bytes
func (s *LSet[T]) Scan(src []byte) error {
	var ls []T
	err := sonic.Unmarshal(src, &ls)
	if err != nil {
		return err
	}
	*s = *NewLSet(ls...)
	return nil
}

// DEBUG
func (s *LSet[T]) Print() {
	fmt.Printf("lset[%d]: %v\n", s.Len(), s.Members())
}

// Compare two lset length and return (*min, *max)
func compareTwoLSet[T comparable](s1 *LSet[T], s2 *LSet[T]) (*LSet[T], *LSet[T]) {
	if s1.Len() <= s2.Len() {
		return s1, s2
	}
	return s2, s1
}
