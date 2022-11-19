package structx

import (
	"fmt"
	"math/rand"
)

const maxLevel = 32
const pFactor = 0.25

type skiplistNode[K, V Value] struct {
	key     K
	value   V
	forward []*skiplistNode[K, V]
}

type Skiplist[K, V Value] struct {
	level int
	len   int
	head  *skiplistNode[K, V]
}

// NewSkipList
func NewSkipList[K, V Value]() *Skiplist[K, V] {
	return &Skiplist[K, V]{
		head: &skiplistNode[K, V]{
			forward: make([]*skiplistNode[K, V], maxLevel),
		},
	}
}

// get random level
func randomLevel() int {
	lv := 1
	for float32(rand.Int31()&0xFFFF) < (pFactor * 0xFFFF) {
		lv++
	}
	if lv < maxLevel {
		return lv
	}
	return maxLevel
}

func (s *Skiplist[K, V]) Len() int {
	return s.len
}

// GetByRank: Get the element by rank
func (s *Skiplist[K, V]) GetByRank(rank int) (k K, v V, err error) {
	p := s.head
	for i := 0; p != nil; i++ {
		if rank == i {
			return p.key, p.value, nil
		}
		p = p.forward[0]
	}

	return k, v, errOutOfBounds(rank)
}

func (s *Skiplist[K, V]) findClosestNode(k K, v V, update []*skiplistNode[K, V]) *skiplistNode[K, V] {
	p := s.head
	for i := s.level - 1; i >= 0; i-- {
		// Find the elem at level[i] that closest to value and key
		// node.value < v || (node.value == v && node.key < k)
		for p.forward[i] != nil && (p.forward[i].value < v || (p.forward[i].value == v && p.forward[i].key < k)) {
			p = p.forward[i]
		}
		update[i] = p
	}
	return p
}

// Add
func (s *Skiplist[K, V]) Add(key K, value V) *skiplistNode[K, V] {
	update := make([]*skiplistNode[K, V], maxLevel)
	for i := range update {
		update[i] = s.head
	}

	s.findClosestNode(key, value, update)

	lv := randomLevel()
	if lv > s.level {
		s.level = lv
	}

	// create node
	newNode := &skiplistNode[K, V]{
		key:     key,
		value:   value,
		forward: make([]*skiplistNode[K, V], lv),
	}

	for i, node := range update[:lv] {
		// Update the state at level[i], pointing the forward of the current element to the new node
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}

	s.len++
	return newNode
}

// Delete
func (s *Skiplist[K, V]) Delete(key K, value V) bool {
	update := make([]*skiplistNode[K, V], maxLevel)

	p := s.findClosestNode(key, value, update)

	p = p.forward[0]
	// if nil or not found
	if p == nil || p.value != value {
		return false
	}

	for i := 0; i < s.level && update[i].forward[i] == p; i++ {
		// Update the state of levek[i] to point forward to the next hop of the deleted node
		update[i].forward[i] = p.forward[i]
	}

	// Update current level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}

	s.len--
	return true
}

// Range
func (s *Skiplist[K, V]) Range(start, end int, f func(key K, value V) bool) {
	if end == -1 {
		end = s.Len()
	}
	p := s.head.forward[0]
	for i := 0; p != nil; i++ {
		// index
		if start <= i && i <= end {
			if f(p.key, p.value) {
				return
			}
		}
		p = p.forward[0]
	}
}

// RevRange
func (s *Skiplist[K, V]) RevRange(start, end int, f func(value V) bool) {
	stack := NewList[V]()
	// push
	s.Range(start, end, func(key K, value V) bool {
		stack.RPush(value)
		return false
	})
	// range
	stack.Range(func(i int, v V) bool {
		return f(v)
	})
}

// RangeByScores
func (s *Skiplist[K, V]) RangeByScores(min, max V, f func(key K, value V) bool) {
	p := s.head.forward[0]
	for p != nil {
		if min <= p.value && p.value <= max {
			if f(p.key, p.value) {
				return
			}
		}
		p = p.forward[0]
	}
}

// DEBUG
func (s *Skiplist[K, V]) Print() {
	fmt.Println("====== start ======")
	s.Range(0, -1, func(key K, value V) bool {
		fmt.Printf("%v -> %v\n", key, value)
		return false
	})
	fmt.Println("======= end =======")
}
