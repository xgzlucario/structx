package structx

import "structx/base"

type List[T base.Value] struct {
	base.Values[T]
}

// NewList: return new List
func NewList[T base.Value]() *List[T] {
	return &List[T]{
		Values: make([]T, 0, 16),
	}
}

// AddToSet: add to the set
func (ls *List[T]) AddToSet(value T) bool {
	if ls.Index(value) < 0 {
		ls.Values = append(ls.Values, value)
		return true
	}
	return false
}

func (ls *List[T]) LPush(value T) {
	ls.RPush(value)
	ls.RShift()
}

func (ls *List[T]) RPush(value T) {
	ls.Values = append(ls.Values, value)
}

func (ls *List[T]) LPop() T {
	ls.LShift()
	return ls.RPop()
}

func (ls *List[T]) RPop() T {
	val := ls.Values[ls.Len()-1]
	ls.Values = ls.Values[:ls.Len()-1]
	return val
}

// Remove: remove value from list
func (ls *List[T]) Remove(value T) bool {
	if i := ls.Index(value); i > 0 {
		ls.Values = append(ls.Values[:i], ls.Values[i+1:]...)
		return true
	}
	return false
}
