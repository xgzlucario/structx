package structx

import "github.com/bytedance/sonic"

type List[T comparable] struct {
	array[T]
}

// NewList: return new List
func NewList[T comparable](values ...T) *List[T] {
	return &List[T]{array: values}
}

// LPush
func (ls *List[T]) LPush(values ...T) *List[T] {
	ls.array = append(values, ls.array...)
	return ls
}

// RPush
func (ls *List[T]) RPush(values ...T) *List[T] {
	ls.array = append(ls.array, values...)
	return ls
}

// LPop
func (ls *List[T]) LPop() T {
	val := ls.array[0]
	ls.array = ls.array[1:]
	return val
}

// Insert
func (ls *List[T]) Insert(i int, value T) *List[T] {
	if i <= 0 {
		ls.LPush(value)

	} else if i >= ls.Len() {
		ls.RPush(value)

	} else {
		ls.array = append(append(ls.array[0:i], value), ls.array[i:]...)
	}
	return ls
}

// RPop
func (ls *List[T]) RPop() T {
	val := ls.array[ls.Len()-1]
	ls.array = ls.array[:ls.Len()-1]
	return val
}

// Remove
func (ls *List[T]) Remove(elem T) bool {
	for i, v := range ls.array {
		if v == elem {
			ls.array = append(ls.array[:i], ls.array[i+1:]...)
			return true
		}
	}
	return false
}

// Marshal: Marshal to bytes
func (s *List[T]) Marshal() ([]byte, error) {
	return sonic.Marshal(s.array)
}

// Scan: Scan from bytes
func (s *List[T]) Scan(src []byte) error {
	return sonic.Unmarshal(src, &s)
}
