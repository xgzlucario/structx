package structx

import "fmt"

type LinkNode[T Comparable] struct {
	val  T
	next *LinkNode[T]
}

// SortList: Sorted Link List
// TODO: Change to ZSet
type SortList[T Comparable] struct {
	head     *LinkNode[T]
	mid      *LinkNode[T]
	tail     *LinkNode[T]
	len      int
	midCount int
}

// NewSortList: Return new SortList with values
func NewSortList[T Comparable](values ...T) *SortList[T] {
	this := &SortList[T]{}
	// init
	for _, value := range values {
		this.Insert(value)
	}
	return this
}

// Insert: Insert and sort value
func (this *SortList[T]) Insert(value T) {
	node := &LinkNode[T]{val: value}

	// head is nil
	if this.head == nil {
		this.head = node
		this.mid = node
		this.tail = node

	} else {
		// less than head
		if value <= this.head.val {
			node.next = this.head
			this.head = node

			// greate than tail
		} else if value >= this.tail.val {
			this.tail.next = node
			this.tail = node

		} else {
			p := findPosition(this, value)
			node.next = p.next
			p.next = node
		}
	}
	this.len++

	// reset mid
	if value < this.mid.val {
		this.midCount++

	} else if this.midCount < this.len/2 {
		this.mid = this.mid.next
	}
}

func (this *SortList[T]) Len() int {
	return this.len
}

// Index: Get element by index
func (this *SortList[T]) Index(index int) T {
	// overflow
	if index >= this.Len() {
		var a T
		return a
	}

	p := this.head
	start := 0

	// start with mid
	if index > this.midCount {
		p = this.mid
		start = this.midCount
	}

	for ; start < index; start++ {
		p = p.next
	}
	return p.val
}

// Values: Return values list
func (this *SortList[T]) Values() []T {
	values := make([]T, this.Len())

	p := this.head
	for i := range values {
		values[i] = p.val
		p = p.next
	}
	return values
}

// find the insert position
func findPosition[T Comparable](this *SortList[T], value T) *LinkNode[T] {
	// compare with mid
	p := this.mid
	if value < this.mid.val {
		p = this.head
	}
	// find
	for ; p.next != nil; p = p.next {
		if value <= p.next.val {
			return p
		}
	}
	return p
}

func (this *SortList[T]) Print() {
	fmt.Printf("SortList len[%d] mid[%d]: ", this.len, this.midCount)
	for p := this.head; p != nil; p = p.next {
		fmt.Printf("%v ", p.val)
	}
	fmt.Println()
}
