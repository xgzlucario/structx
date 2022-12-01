package structx

import (
	"sort"

	"github.com/bytedance/sonic"
)

type List[T comparable] struct {
	array[T]
	order bool
	less  func(T, T) bool // the input params is elements
}

// NewList: return new List
func NewList[T comparable](values ...T) *List[T] {
	return &List[T]{array: values, order: true}
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

// SetOrder
func (ls *List[T]) SetOrder(order bool) *List[T] {
	ls.order = order
	return ls
}

/*
SetLess: Set Less to make list sortable. The input params is elements.

exp:

	ls.SetLess(func(i, j int) bool {
		return i < j
	})

or:

	type A struct {
		Name string
		Score int
	}
	ls.SetLess(func(i, j A) bool {
		return i.Score < j.Score
	})
*/
func (ls *List[T]) SetLess(f func(T, T) bool) *List[T] {
	ls.less = f
	return ls
}

func (ls *List[T]) Less(i, j int) bool {
	if ls.order {
		// arr[i] < arr[j]
		return ls.less(ls.array[i], ls.array[j])
	}
	// arr[j] < arr[i]
	return ls.less(ls.array[j], ls.array[i])
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
	sort.Sort(ls)
	return ls
}

// IsSorted: Should SetLess First
func (ls *List[T]) IsSorted() bool {
	return sort.IsSorted(ls)
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
