package structx

import "sort"

type NumList[T Number] struct {
	*List[T]
	order bool // sort
}

// NewNumList
func NewNumList[T Number](values ...T) *NumList[T] {
	return &NumList[T]{
		List:  NewList(values...),
		order: true,
	}
}

// Max
func (ls *NumList[T]) Max() T {
	var max = ls.array[0]
	for _, v := range ls.array {
		if v > max {
			max = v
		}
	}
	return max
}

// Min
func (ls *NumList[T]) Min() T {
	var min = ls.array[0]
	for _, v := range ls.array {
		if v < min {
			min = v
		}
	}
	return min
}

// Sum
func (ls *NumList[T]) Sum() T {
	var sum T
	for _, v := range ls.array {
		sum += v
	}
	return sum
}

// Mean
func (ls *NumList[T]) Mean() float64 {
	var sum T
	for _, v := range ls.array {
		sum += v
	}
	return float64(sum) / float64(ls.Len())
}

// Median
func (ls *NumList[T]) Median() float64 {
	var sum T
	for _, v := range ls.array {
		sum += v
	}
	return float64(sum) / float64(ls.Len())
}

// Less
func (ls *NumList[T]) Less(i, j int) bool {
	if ls.order {
		return ls.array[i] < ls.array[j]
	}
	return ls.array[i] > ls.array[j]
}

// Sort
func (ls *NumList[T]) Sort() {
	sort.Sort(ls)
}

// SetOrder
func (ls *NumList[T]) SetOrder(order bool) {
	ls.order = order
}

// IsSorted
func (ls *NumList[T]) IsSorted() bool {
	return sort.IsSorted(ls)
}
