package structx

import "fmt"

type linkNode[K any, T Value] struct {
	key   K
	value T
	next  *linkNode[K, T]
}

// SortList: Sorted Link List
type SortList[K any, T Value] struct {
	head     *linkNode[K, T]
	mid      *linkNode[K, T]
	tail     *linkNode[K, T]
	len      int
	midCount int
}

// NewSortList: Return new SortList with values
func NewSortList[K any, T Value]() *SortList[K, T] {
	return &SortList[K, T]{}
}

// Insert: Insert and sort value
func (this *SortList[K, T]) Insert(value T, key ...K) {
	node := &linkNode[K, T]{value: value}
	if len(key) > 0 {
		node.key = key[0]
	}

	// head is nil
	if this.head == nil {
		this.head = node
		this.mid = node
		this.tail = node

	} else {
		// less than head
		if value <= this.head.value {
			node.next = this.head
			this.head = node

			// greate than tail
		} else if value >= this.tail.value {
			this.tail.next = node
			this.tail = node

		} else {
			p := this.findPosition(value)
			node.next = p.next
			p.next = node
		}
	}
	this.len++

	// reset mid
	if value < this.mid.value {
		this.midCount++

	} else if this.midCount < this.len/2 {
		this.mid = this.mid.next
		this.midCount++
	}
}

// Delete: Delete the first node of value
func (this *SortList[K, T]) Delete(value T) *linkNode[K, T] {
	for p := this.head; p.next != nil; p = p.next {
		// find
		if p.next.value == value {
			// delete
			node := p.next
			p.next = p.next.next
			return node
		}
	}
	return nil
}

func (this *SortList[K, T]) Empty() bool {
	return this.head == nil
}

func (this *SortList[K, T]) Len() int {
	return this.len
}

// Index: Get element by index
func (this *SortList[K, T]) Index(index int) *linkNode[K, T] {
	// overflow
	if index >= this.Len() {
		return nil
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
	return p
}

// Values: Return values list
func (this *SortList[K, T]) Values() []T {
	values := make([]T, this.Len())

	p := this.head
	for i := range values {
		values[i] = p.value
		p = p.next
	}
	return values
}

// find the insert position
func (this *SortList[K, T]) findPosition(value T) *linkNode[K, T] {
	// compare with mid
	p := this.mid
	if value < this.mid.value {
		p = this.head
	}
	// find
	for ; p.next != nil; p = p.next {
		if value <= p.next.value {
			return p
		}
	}
	return p
}

func (this *SortList[K, T]) Print() {
	fmt.Printf("SortList len[%d] mid[%d]: ", this.len, this.midCount)
	for p := this.head; p != nil; p = p.next {
		fmt.Printf("%v ", p.value)
	}
	fmt.Println()
}
