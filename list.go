package structx

import (
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
		ls.array = append(append(ls.array[0:i], values...), ls.array[i:]...)
	}
	return ls
}

// LPop
func (ls *List[T]) LPop() (v T, ok bool) {
	if ls.Len() == 0 {
		return
	}
	val := ls.array[0]
	ls.array = ls.array[1:]
	return val, true
}

// RPop
func (ls *List[T]) RPop() (v T, ok bool) {
	if ls.Len() == 0 {
		return
	}
	val := ls.array[ls.Len()-1]
	ls.array = ls.array[:ls.Len()-1]
	return val, true
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
	return marshalBin(s.array)
}

// Unmarshal
func (s *List[T]) Unmarshal(src []byte) error {
	return unmarshalBin(src, &s.array)
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
