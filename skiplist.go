package structx

import "math/rand"

const maxLevel = 32
const pFactor = 0.25

type skiplistNode[T Value] struct {
	val     T
	forward []*skiplistNode[T]
}

type Skiplist[T Value] struct {
	head  *skiplistNode[T]
	level int
}

// NewSkipList
func NewSkipList[T Value]() *Skiplist[T] {
	return &Skiplist[T]{
		head: &skiplistNode[T]{
			forward: make([]*skiplistNode[T], maxLevel),
		},
	}
}

func (Skiplist[T]) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

// Search
func (s *Skiplist[T]) Search(target T) bool {
	p := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for p.forward[i] != nil && p.forward[i].val < target {
			p = p.forward[i]
		}
	}

	p = p.forward[0]
	// 检测当前元素的值是否等于 target
	return p != nil && p.val == target
}

// Add
func (s *Skiplist[T]) Add(value T) {
	update := make([]*skiplistNode[T], maxLevel)
	for i := range update {
		update[i] = s.head
	}

	p := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 value 的元素
		for p.forward[i] != nil && p.forward[i].val < value {
			p = p.forward[i]
		}
		update[i] = p
	}

	lv := s.randomLevel()
	if lv > s.level {
		s.level = lv
	}

	newNode := &skiplistNode[T]{
		val:     value,
		forward: make([]*skiplistNode[T], lv),
	}

	for i, node := range update[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

// Delete
func (s *Skiplist[T]) Delete(value T) bool {
	update := make([]*skiplistNode[T], maxLevel)
	p := s.head

	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for p.forward[i] != nil && p.forward[i].val < value {
			p = p.forward[i]
		}
		update[i] = p
	}

	p = p.forward[0]
	// 如果值不存在则返回 false
	if p == nil || p.val != value {
		return false
	}

	for i := 0; i < s.level && update[i].forward[i] == p; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = p.forward[i]
	}

	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}
