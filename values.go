package structx

import (
	"fmt"
	"sort"
)

type Value interface {
	string | float64 | float32 | int64 | int32 | int | uint | byte
}

type Values[T Value] []T

func (s Values[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Values[T]) Len() int {
	return len(s)
}

func (s Values[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

// Sort: use sort
func (s Values[T]) Sort() {
	sort.Sort(s)
}

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

// Index: return the first index of elem
func (s Values[T]) Index(value T) int {
	for i, v := range s {
		if v == value {
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

// Move: Move an element to any position
func (s Values[T]) Move(old, new int) {
	// TODO
}

func (s Values[T]) Reverse() {
	l, r := 0, s.Len()-1
	for l < r {
		s.Swap(l, r)
		l++
		r--
	}
}

func (this Values[T]) Max() T {
	max := this[0]
	for _, s := range this {
		if s > max {
			max = s
		}
	}
	return max
}

func (this Values[T]) Min() T {
	min := this[0]
	for _, s := range this {
		if s < min {
			min = s
		}
	}
	return min
}

func (this Values[T]) Print() {
	fmt.Printf("values: %v\n", this)
}
