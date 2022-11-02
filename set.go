package structx

// Set: map + list (not thread safe)
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
		s.ls.Bottom(i)
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

func (s *Set[T]) Copy(dst *Set[T]) {
	// copy list
	copy(dst.ls.Values, s.ls.Values)
	// copy map
	s.m.Range(func(k T, v struct{}) bool {
		dst.m.Store(k, v)
		return false
	})
}

func (s *Set[T]) Intersection(t Set[T]) {
	t.Range(func(k T) bool {
		if !s.Exist(k) {
			s.Remove(k)
		}
		return false
	})
}

func (s *Set[T]) Sort() {
	s.ls.Sort()
}

func (s *Set[T]) Reverse() {
	s.ls.Reverse()
}

func (s *Set[T]) Values() Values[T] {
	return s.ls.Values
}
