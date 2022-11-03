package structx

import "fmt"

/*
LSet (ListSet): map + list (not thread safe)
ListSet has a significant performance improvement over MapSet
in the Range Union Intersect function
*/
type LSet[T Value] struct {
	m  Map[T, struct{}]
	ls *List[T]
}

func NewLSet[T Value](values ...T) *LSet[T] {
	// make map
	m := make(Map[T, struct{}], len(values))
	for _, v := range values {
		m.Store(v, struct{}{})
	}
	return &LSet[T]{
		m: m, ls: NewList(values...),
	}
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

// make sure that key is not exist!
func (s *LSet[T]) add(key T) {
	s.ls.RPush(key)
	s.m.Store(key, struct{}{})
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

// make sure that key is exist!
func (s *LSet[T]) remove(key T) {
	s.m.Delete(key)
	s.ls.RemoveElem(key)
}

func (s *LSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s *LSet[T]) Range(f func(k T)) {
	for _, value := range s.ls.Values {
		f(value)
	}
}

func (s *LSet[T]) Copy() *LSet[T] {
	arr := make([]T, s.ls.Len())
	copy(arr, s.Values())
	return NewLSet(arr...)
}

// Union
func (this *LSet[T]) Union(t *LSet[T]) *LSet[T] {
	min, max := compareTwoLSet(this, t)
	// should copy max lset
	max = max.Copy()
	min.Range(func(k T) {
		max.Add(k)
	})
	return max
}

// Intersect
func (this *LSet[T]) Intersect(t *LSet[T]) *LSet[T] {
	min, max := compareTwoLSet(this, t)
	// should copy min lset
	min = min.Copy()
	min.Range(func(k T) {
		if !max.Exist(k) {
			min.remove(k)
		}
	})
	return min
}

// Difference
func (this *LSet[T]) Difference(t *LSet[T]) *LSet[T] {
	newSet := NewLSet[T]()
	this.Range(func(k T) {
		if !t.Exist(k) {
			newSet.add(k)
		}
	})
	t.Range(func(k T) {
		if !this.Exist(k) {
			newSet.add(k)
		}
	})
	return newSet
}

func (s *LSet[T]) Sort() {
	s.ls.Sort()
}

func (s *LSet[T]) Reverse() {
	s.ls.Reverse()
}

func (s *LSet[T]) Len() int {
	return s.ls.Len()
}

func (s *LSet[T]) Min() T {
	return s.ls.Min()
}

func (s *LSet[T]) Max() T {
	return s.ls.Max()
}

func (s *LSet[T]) Values() Values[T] {
	return s.ls.Values
}

func (s *LSet[T]) Print() {
	fmt.Println("lset:", s.Values())
}

// Compare two lset length and return (*min, *max)
func compareTwoLSet[T Value](s1 *LSet[T], s2 *LSet[T]) (*LSet[T], *LSet[T]) {
	if s1.Len() <= s2.Len() {
		return s1, s2
	}
	return s2, s1
}
