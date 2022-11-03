package structx

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

func (s *LSet[T]) Add(key T) bool {
	_, ok := s.m[key]
	if ok {
		return false
	}
	s.ls.RPush(key)
	s.m.Store(key, struct{}{})
	return true
}

func (s LSet[T]) Remove(key T) bool {
	_, ok := s.m[key]
	if ok {
		// remove map
		s.m.Delete(key)
		// remove list
		s.ls.RemoveElem(key)
		return true
	}
	return false
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
func (s *LSet[T]) Union(t *LSet[T]) *LSet[T] {
	newLSet := s.Copy()
	t.Range(func(k T) {
		newLSet.Add(k)
	})
	return newLSet
}

// Intersect
func (s *LSet[T]) Intersect(t *LSet[T]) *LSet[T] {
	newLSet := s.Copy()
	t.Range(func(k T) {
		if !newLSet.Exist(k) {
			newLSet.Remove(k)
		}
	})
	return newLSet
}

func (s *LSet[T]) Sort() {
	s.ls.Sort()
}

func (s *LSet[T]) Reverse() {
	s.ls.Reverse()
}

func (s *LSet[T]) Values() Values[T] {
	return s.ls.Values
}
