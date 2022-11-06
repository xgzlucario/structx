package structx

import "fmt"

type ZSetNode[K comparable, T Comparable] struct {
	key   K
	value T
	next  *ZSetNode[K, T]
}

// ZSetV2: map + skipList
type ZSetV2[K comparable, T Comparable] struct {
	head     *ZSetNode[K, T]
	mid      *ZSetNode[K, T]
	tail     *ZSetNode[K, T]
	m        Map[K, *ZSetNode[K, T]]
	midCount int
}

// NewZSet: Return new ZSet with values
func NewZSet[K comparable, T Comparable]() *ZSetV2[K, T] {
	return &ZSetV2[K, T]{
		m: Map[K, *ZSetNode[K, T]]{},
	}
}

// Insert: Insert and sort value
func (this *ZSetV2[K, T]) Insert(key K, value T) {
	node := &ZSetNode[K, T]{
		key: key, value: value,
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

	// reset mid
	if value < this.mid.value {
		this.midCount++

	} else if this.midCount < this.Len()/2 {
		this.mid = this.mid.next
	}
}

// Delete: Delete value
func (this *ZSetV2[K, T]) Delete(value T) {

}

func (this *ZSetV2[K, T]) Len() int {
	return len(this.m)
}

// Index: Get element by index
func (this *ZSetV2[K, T]) Index(index int) T {
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
	return p.value
}

// Values: Return values list
func (this *ZSetV2[K, T]) Values() []T {
	values := make([]T, this.Len())

	p := this.head
	for i := range values {
		values[i] = p.value
		p = p.next
	}
	return values
}

// find the insert position
func (this *ZSetV2[K, T]) findPosition(value T) *ZSetNode[K, T] {
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

func (this *ZSetV2[K, T]) Print() {
	fmt.Printf("SortList len[%d] mid[%d]: ", this.Len(), this.midCount)
	for p := this.head; p != nil; p = p.next {
		fmt.Printf("%v ", p.value)
	}
	fmt.Println()
}
