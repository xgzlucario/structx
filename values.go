package structx

import (
	"fmt"
)

type Comparable interface {
	string | float32 | float64 | int64 | int32 | int | uint | byte
}

type Values[T comparable] []T

func (s Values[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Values[T]) Len() int {
	return len(s)
}

// func (s Values[T]) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s Values[T]) Sort() {
// 	sort.Sort(s)
// }

// Top: move value to the top
func (s Values[T]) Top(i int) {
	for j := i; j > 0; j-- {
		s.Swap(j, j-1)
	}
}

// Bottom: move value to the bottom
func (s Values[T]) Bottom(i int) {
	for j := i; j < s.Len()-1; j++ {
		s.Swap(j, j+1)
	}
}

// Index: return the element of index
func (this Values[T]) Index(index int) T {
	return this[index]
}

// Find: return the index of element
func (this Values[T]) Find(elem T) int {
	for i, v := range this {
		if v == elem {
			return i
		}
	}
	return -1
}

// LShift: Shift all elements of the array left
// exp: [1, 2, 3] => [2, 3, 1]
func (s Values[T]) LShift() {
	s.Bottom(0)
}

// RShift: Shift all elements of the array right
// exp: [1, 2, 3] => [3, 1, 2]
func (s Values[T]) RShift() {
	s.Top(s.Len() - 1)
}

func (s Values[T]) Reverse() {
	l, r := 0, s.Len()-1
	for l < r {
		s.Swap(l, r)
		l++
		r--
	}
}

func (this Values[T]) Range(f func(i int, v T)) {
	for i, v := range this {
		f(i, v)
	}
}

func (this Values[T]) Print() {
	fmt.Println("values:", this)
}
