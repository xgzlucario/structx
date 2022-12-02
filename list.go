package structx

import (
	"golang.org/x/exp/slices"

	"github.com/bytedance/sonic"
)

type List[T comparable] struct {
	array[T]
	order func(T, T) bool // Sort, IsSorted Used
	less  func(T, T) bool // Max, Min Used
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

// LPop
func (ls *List[T]) LPop() T {
	val := ls.array[0]
	ls.array = ls.array[1:]
	return val
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

// SetOrder: Make list sortable. The input params is elements.
func (ls *List[T]) SetOrder(f func(T, T) bool) *List[T] {
	ls.order = f
	return ls
}

// SetLess: Make list sortable. The input params is elements.
func (ls *List[T]) SetLess(f func(T, T) bool) *List[T] {
	ls.less = f
	return ls
}

// Max: Should SetLess First
func (ls *List[T]) Max() T {
	max := ls.array[0]
	for _, v := range ls.array {
		if ls.less(max, v) {
			max = v
		}
	}
	return max
}

// Min: Should SetLess First
func (ls *List[T]) Min() T {
	min := ls.array[0]
	for _, v := range ls.array {
		if ls.less(v, min) {
			min = v
		}
	}
	return min
}

// Sort: Should SetLess First
func (ls *List[T]) Sort() *List[T] {
	ls.checkOrder()
	slices.SortFunc(ls.array, ls.order)
	return ls
}

// IsSorted: Should SetLess First
func (ls *List[T]) IsSorted() bool {
	ls.checkOrder()
	return slices.IsSortedFunc(ls.array, ls.order)
}

// check if order is nil
func (ls *List[T]) checkOrder() {
	// default ascending, please set less first
	if ls.order == nil {
		if ls.less == nil {
			panic("Please use SetLess() to init less first")
		}
		ls.order = ls.less
	}
}

// Marshal: Marshal to bytes
func (s *List[T]) Marshal() ([]byte, error) {
	return sonic.Marshal(s.array)
}

// Scan: Scan from bytes
func (s *List[T]) Scan(src []byte) error {
	return sonic.Unmarshal(src, &s)
}

// Print
func (s *List[T]) Print() {
	s.array.Print()
}
