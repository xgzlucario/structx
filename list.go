package structx

import "github.com/bytedance/sonic"

type List[T comparable] struct {
	Array[T]
}

// NewList: return new List
func NewList[T comparable](values ...T) *List[T] {
	return &List[T]{Array: values}
}

// LPush
func (ls *List[T]) LPush(values ...T) {
	ls.Array = append(values, ls.Array...)
}

// RPush
func (ls *List[T]) RPush(values ...T) {
	ls.Array = append(ls.Array, values...)
}

// LPop
func (ls *List[T]) LPop() T {
	val := ls.Array[0]
	ls.Array = ls.Array[1:]
	return val
}

// Insert
func (ls *List[T]) Insert(i int, value T) {
	if i <= 0 {
		ls.LPush(value)

	} else if i >= ls.Len() {
		ls.RPush(value)

	} else {
		ls.Array = append(append(ls.Array[0:i], value), ls.Array[i:]...)
	}
}

// RPop
func (ls *List[T]) RPop() T {
	val := ls.Array[ls.Len()-1]
	ls.Array = ls.Array[:ls.Len()-1]
	return val
}

// Remove
func (ls *List[T]) Remove(elem T) bool {
	for i, v := range ls.Array {
		if v == elem {
			ls.Array = append(ls.Array[:i], ls.Array[i+1:]...)
			return true
		}
	}
	return false
}

// Marshal: Marshal to bytes
func (s *List[T]) Marshal() ([]byte, error) {
	return sonic.Marshal(s.Array)
}

// Scan: Scan from bytes
func (s *List[T]) Scan(src []byte) error {
	return sonic.Unmarshal(src, &s)
}
