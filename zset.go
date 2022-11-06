package structx

import "fmt"

type zsetNode[K, V Value] struct {
	key   K
	value V
}

type ZSet[K, V Value] struct {
	zsl *Skiplist[V]
	m   Map[K, *zsetNode[K, V]]
}

// NewZSet
func NewZSet[K, V Value]() *ZSet[K, V] {
	return &ZSet[K, V]{
		zsl: NewSkipList[V](),
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
		z.zsl.Delete(n.value)
		z.zsl.Add(n.value)

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
	return p.val
}

func (z *ZSet[K, V]) Len() int {
	return len(z.m)
}

func (z *ZSet[K, V]) insertNode(key K, value V) {
	z.m[key] = &zsetNode[K, V]{
		key:   key,
		value: value,
	}
	z.zsl.Add(value)
}

func (z *ZSet[K, V]) deleteNode(key K, value V) {
	delete(z.m, key)
	z.zsl.Delete(value)
}

func (z *ZSet[K, V]) Print() {
	for _, p := range z.zsl.head.forward {
		fmt.Println(p.val)
	}
}
