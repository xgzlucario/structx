package structx

import (
	"fmt"
)

type Array[T comparable] []T

func (s Array[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Array[T]) Len() int {
	return len(s)
}

// Top: move value to the top
func (s Array[T]) Top(i int) {
	for j := i; j > 0; j-- {
		s.Swap(j, j-1)
	}
}

// Bottom: move value to the bottom
func (s Array[T]) Bottom(i int) {
	for j := i; j < s.Len()-1; j++ {
		s.Swap(j, j+1)
	}
}

// Index: return the element of index
func (this Array[T]) Index(index int) T {
	return this[index]
}

// Find: return the index of element
func (this Array[T]) Find(elem T) int {
	for i, v := range this {
		if v == elem {
			return i
		}
	}
	return -1
}

// LShift: Shift all elements of the array left
// exp: [1, 2, 3] => [2, 3, 1]
func (s Array[T]) LShift() {
	s.Bottom(0)
}

// RShift: Shift all elements of the array right
// exp: [1, 2, 3] => [3, 1, 2]
func (s Array[T]) RShift() {
	s.Top(s.Len() - 1)
}

func (s Array[T]) Reverse() {
	l, r := 0, s.Len()-1
	for l < r {
		s.Swap(l, r)
		l++
		r--
	}
}

func (this Array[T]) Range(f func(int, T) bool) {
	for i, v := range this {
		if f(i, v) {
			return
		}
	}
}

// DEBUG
func (this Array[T]) Print() {
	fmt.Println("Array:", this)
}
