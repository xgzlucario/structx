package structx

import (
	"math/rand"
)

const maxLevel = 32
const pFactor = 0.25

type skiplistNode[K comparable, V Value] struct {
	key     K
	value   V
	forward []*skiplistNode[K, V]
}

type Skiplist[K comparable, V Value] struct {
	head  *skiplistNode[K, V]
	level int
	len   int
}

// NewSkipList
func NewSkipList[K comparable, V Value]() *Skiplist[K, V] {
	return &Skiplist[K, V]{
		head: &skiplistNode[K, V]{
			forward: make([]*skiplistNode[K, V], maxLevel),
		},
	}
}

func (Skiplist[K, V]) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (s *Skiplist[K, V]) Len() int {
	return s.len
}

// Search
func (s *Skiplist[K, V]) Search(value V, key ...K) bool {
	p := s.head
	for i := s.level - 1; i >= 0; i-- {
		// Find the element at level[i] that is less than and closest to value
		for p.forward[i] != nil && p.forward[i].value < value {
			p = p.forward[i]
		}
	}

	p = p.forward[0]

	// Check if the value of the current element is equal to value
	if len(key) > 0 {
		return p != nil && p.value == value && p.key == key[0]
	}
	return p != nil && p.value == value
}

// Add
func (s *Skiplist[K, V]) Add(value V, key ...K) *skiplistNode[K, V] {
	update := make([]*skiplistNode[K, V], maxLevel)
	for i := range update {
		update[i] = s.head
	}

	p := s.head
	for i := s.level - 1; i >= 0; i-- {
		// Find the element at level[i] that is less than and closest to value
		for p.forward[i] != nil && p.forward[i].value < value {
			p = p.forward[i]
		}
		update[i] = p
	}

	lv := s.randomLevel()
	if lv > s.level {
		s.level = lv
	}

	// create node
	newNode := &skiplistNode[K, V]{
		value:   value,
		forward: make([]*skiplistNode[K, V], lv),
	}
	if len(key) > 0 {
		newNode.key = key[0]
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
func (s *Skiplist[K, V]) Delete(value V, key ...K) bool {
	update := make([]*skiplistNode[K, V], maxLevel)
	p := s.head

	for i := s.level - 1; i >= 0; i-- {
		// Find the element at level[i] that is less than and closest to value
		for p.forward[i] != nil && p.forward[i].value < value {
			p = p.forward[i]
		}
		update[i] = p
	}

	p = p.forward[0]
	// if nil or not found
	if p == nil || p.value != value {
		return false
	}

	// key not equal
	if len(key) > 0 && key[0] != p.key {
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
func (s *Skiplist[K, V]) Range(start, end int, f func(key K, value V)) {
	if end == -1 {
		end = s.Len()
	}
	var (
		now  int
		read func(p *skiplistNode[K, V])
	)
	read = func(p *skiplistNode[K, V]) {
		if p != nil {
			// index
			if start <= now && now <= end {
				f(p.key, p.value)
			}
			now++
			// recursive
			read(p.forward[0])
		}
	}
	read(s.head.forward[0])
}

// RevRange
func (s *Skiplist[K, V]) RevRange(start, end int, f func(value V)) {
	stack := NewList[V]()
	// push
	s.Range(start, end, func(key K, value V) {
		stack.RPush(value)
	})
	// range
	stack.Range(func(i int, v V) {
		f(v)
	})
}

// RangeByScores
func (s *Skiplist[K, V]) RangeByScores(min, max V, f func(key K, value V)) {
	var read func(p *skiplistNode[K, V])

	read = func(p *skiplistNode[K, V]) {
		if p != nil {
			if min <= p.value && p.value <= max {
				f(p.key, p.value)
			}
			// recursive
			read(p.forward[0])
		}
	}
	read(s.head.forward[0])
}

// GetByRank
func (s *Skiplist[K, V]) GetByRank(index int, f func(key K, value V)) {
	s.Range(index, index, f)
}
