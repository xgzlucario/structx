package structx

import "fmt"

type zsetNode[K comparable, V Value] struct {
	key   K
	value V
}

type ZSet[K comparable, V Value] struct {
	zsl *Skiplist[K, V]
	m   Map[K, *zsetNode[K, V]]
}

// NewZSet
func NewZSet[K comparable, V Value]() *ZSet[K, V] {
	return &ZSet[K, V]{
		zsl: NewSkipList[K, V](),
		m:   Map[K, *zsetNode[K, V]]{},
	}
}

// Set
func (z *ZSet[K, V]) Set(key K, value V) bool {
	n, ok := z.m[key]

	// value not change
	if value == n.value {
		return false
	}
	if ok {
		n.value = value
		z.zsl.Delete(n.value, key)
		z.zsl.Add(n.value, key)

	} else {
		z.insertNode(key, value)
	}
	return true
}

// Incr
func (z *ZSet[K, V]) Incr(key K, value V) V {
	n, ok := z.m[key]
	// not exist
	if !ok {
		z.insertNode(key, value)
		return value
	}
	// exist
	z.zsl.Delete(n.value)
	n.value += value
	z.zsl.Add(n.value)

	return n.value
}

// Delete
func (z *ZSet[K, V]) Delete(key K) bool {
	n, ok := z.m[key]
	if ok {
		z.deleteNode(n.key, n.value)
	}
	return ok
}

// GetByRank
func (z *ZSet[K, V]) GetByRank(rank int) V {
	p := z.zsl.head.forward[rank]
	return p.value
}

func (z *ZSet[K, V]) Len() int {
	return len(z.m)
}

func (z *ZSet[K, V]) insertNode(key K, value V) *skiplistNode[K, V] {
	z.m[key] = &zsetNode[K, V]{
		key:   key,
		value: value,
	}
	// add node
	return z.zsl.Add(value, key)
}

func (z *ZSet[K, V]) deleteNode(key K, value V) {
	delete(z.m, key)
	// delete node
	z.zsl.Delete(value, key)
}

func (z *ZSet[K, V]) Print() {
	for _, p := range z.zsl.head.forward {
		fmt.Println(p.value)
	}
}
