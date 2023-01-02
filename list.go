package structx

import (
	"fmt"

	"golang.org/x/exp/slices"
)

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

// Insert
func (ls *List[T]) Insert(i int, values ...T) *List[T] {
	if i <= 0 {
		ls.LPush(values...)
	} else if i >= ls.Len() {
		ls.RPush(values...)
	} else {
		ls.array = slices.Insert(ls.array, i, values...)
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

// RemoveElem
func (ls *List[T]) RemoveElem(elem T) error {
	for i, v := range ls.array {
		if v == elem {
			ls.array = slices.Delete(ls.array, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("elem[%v] not exist", elem)
}

// Delete
func (ls *List[T]) Delete(i, j int) *List[T] {
	ls.array = slices.Delete(ls.array, i, j)
	return ls
}

// Max: Input param is Less function
func (ls *List[T]) Max(less func(T, T) bool) T {
	max := ls.array[0]
	for _, v := range ls.array {
		if less(max, v) {
			max = v
		}
	}
	return max
}

// Min: Input param is Less function
func (ls *List[T]) Min(less func(T, T) bool) T {
	min := ls.array[0]
	for _, v := range ls.array {
		if less(v, min) {
			min = v
		}
	}
	return min
}

// Sum
func (ls *List[T]) Sum(toNumber func(T) float64) float64 {
	var sum float64
	for _, v := range ls.array {
		sum += toNumber(v)
	}
	return sum
}

// Mean
func (ls *List[T]) Mean(toNumber func(T) float64) float64 {
	return ls.Sum(toNumber) / float64(ls.Len())
}

// Sort: Input param is Order function
func (ls *List[T]) Sort(order func(T, T) bool) *List[T] {
	slices.SortFunc(ls.array, order)
	return ls
}

// IsSorted: Input param is Order function
func (ls *List[T]) IsSorted(order func(T, T) bool) bool {
	return slices.IsSortedFunc(ls.array, order)
}

// Marshal
func (s *List[T]) Marshal() ([]byte, error) {
	return marshalJSON(s.array)
}

// Unmarshal
func (s *List[T]) Unmarshal(src []byte) error {
	return unmarshalJSON(src, &s.array)
}

// Values
func (s *LSet[T]) Values() array[T] {
	return s.array
}

// Print
func (s *List[T]) Print() *List[T] {
	s.array.Print()
	return s
}
