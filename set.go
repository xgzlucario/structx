package structx

// Set: map + list
type Set[T Value] struct {
	m  Map[T, struct{}]
	ls *List[T]
}

func NewSet[T Value]() Set[T] {
	return Set[T]{
		m:  NewMap[T, struct{}](),
		ls: NewList[T](),
	}
}

func (s *Set[T]) Add(key T) bool {
	_, ok := s.m[key]
	if ok {
		return false
	}
	s.ls.RPush(key)
	s.m.Store(key, struct{}{})
	return false
}

func (s Set[T]) Remove(key T) bool {
	_, ok := s.m[key]
	if ok {
		// remove map
		s.m.Delete(key)
		// remove list
		i := s.ls.Index(key)
		s.ls.Swap(i, s.ls.Len()-1)
		s.ls.RPop()
		return true
	}
	return false
}

func (s *Set[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set[T]) Range(f func(k T) bool) {
	s.ls.Range(f)
}

func (s *Set[T]) Union(t Set[T]) {
	t.Range(func(k T) bool {
		s.Add(k)
		return false
	})
}

func (s *Set[T]) Intersection(t Set[T]) {
	t.Range(func(k T) bool {
		s.Add(k)
		return false
	})
}

func (s *Set[T]) Sort() {
	s.ls.Sort()
}

func (s *Set[T]) All() Values[T] {
	return s.ls.Values
}
